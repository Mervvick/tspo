package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

type Profile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	LastSeen string `json:"last_seen"`
}

var redisClient *redis.Client
var ctx = context.Background()

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// Имитация "медленного" получения данных
func getProfileFromSlowSource(id string) (Profile, error) {
	// Имитируем задержку
	time.Sleep(2 * time.Second)

	// Имитируем получение данных из базы
	profile := Profile{
		ID:       id,
		Name:     fmt.Sprintf("User %s", id),
		Email:    fmt.Sprintf("user%s@example.com", id),
		Age:      25 + (len(id) % 30), // Генерация случайного возраста
		LastSeen: time.Now().Format(time.RFC3339),
	}

	return profile, nil
}

func getProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	cacheKey := fmt.Sprintf("profile:%s", id)

	// Пытаемся получить данные из кэша
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Данные найдены в кэше
		log.Printf("Cache hit for profile: %s", id)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Cache", "HIT")
		w.Write([]byte(cachedData))
		return
	}

	// Если данных нет в кэше, получаем из "медленного источника"
	log.Printf("Cache miss for profile: %s, fetching from source", id)
	profile, err := getProfileFromSlowSource(id)
	if err != nil {
		http.Error(w, "Error getting profile", http.StatusInternalServerError)
		return
	}

	// Сериализуем профиль
	profileJSON, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, "Error serializing profile", http.StatusInternalServerError)
		return
	}

	// Сохраняем в кэш на 10 минут
	err = redisClient.Set(ctx, cacheKey, profileJSON, 10*time.Minute).Err()
	if err != nil {
		log.Printf("Error caching profile: %v", err)
	}

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Cache", "MISS")
	w.Write(profileJSON)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/profile/{id}", getProfileHandler)

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
