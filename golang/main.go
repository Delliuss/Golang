// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	"os"
// 	"strings"
// )

// // Структура для персонажа
// type Character struct {
// 	Name        string // Имя героя
// 	Description string // Краткое описание героя
// 	ImageURL    string // Фото персонажа
// 	ReleaseDate string // Дата выхода героя
// 	Role        string // Роль в игре
// 	Class       string // Тип урона персонажа
// 	Price       int    // Цена персонажа в игре
// }

// // Структура для пользователя
// type User struct {
// 	Username string // Имя пользователя
// 	Password string // Пароль пользователя
// 	Email    string //
// }

// // Список персонажей
// var characters []Character

// // Функция для загрузки персонажей из файла
// func loadCharacters() error {
// 	file, err := os.Open("characters.json")
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&characters); err != nil {
// 		return err
// 	}
// 	return nil
// }

// var users []User

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	// Заголовок
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	// Получаем значение параметра поиска
// 	searchQuery := r.URL.Query().Get("search")
// 	var filteredCharacters []Character

// 	// Фильтруем персонажей по запросу
// 	for _, character := range characters {
// 		if searchQuery == "" || strings.Contains(strings.ToLower(character.Name), strings.ToLower(searchQuery)) {
// 			filteredCharacters = append(filteredCharacters, character)
// 		}
// 	}

// 	// Создание структуры для передачи данных в шаблон
// 	data := struct {
// 		Characters  []Character
// 		SearchQuery string
// 	}{
// 		Characters:  filteredCharacters,
// 		SearchQuery: searchQuery,
// 	}

// 	// Загружаем HTML-шаблон из файла
// 	tmpl, err := template.ParseFiles("html_go.html")
// 	if err != nil {
// 		http.Error(w, "Ошибка при загрузке шаблона", http.StatusInternalServerError)
// 		return
// 	}

// 	// Выполнение шаблона и передача данных
// 	if err := tmpl.Execute(w, data); err != nil {
// 		http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
// 	}
// }

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		// Получаем данные из формы
// 		username := r.FormValue("username")
// 		password := r.FormValue("password")

// 		// Проверяем, существует ли уже пользователь с таким именем
// 		for _, user := range users {
// 			if user.Username == username {
// 				http.Error(w, "Пользователь с таким именем уже существует", http.StatusConflict)
// 				return
// 			}
// 		}

// 		// Добавляем нового пользователя в список
// 		users = append(users, User{Username: username, Password: password})

// 		// Перенаправляем на главную страницу после успешной регистрации
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	// Обработка метода GET для отображения страницы регистрации
// 	if r.Method == http.MethodGet {
// 		// Загружаем HTML-шаблон для регистрации
// 		tmpl, err := template.ParseFiles("register.html")
// 		if err != nil {
// 			http.Error(w, "Ошибка при загрузке шаблона", http.StatusInternalServerError)
// 			return
// 		}

// 		// Выполнение шаблона и передача данных (если нужно)
// 		if err := tmpl.Execute(w, nil); err != nil {
// 			http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	http.Error(w, "Не разрешен", http.StatusMethodNotAllowed)
// }

// func main() {
// 	// Загружаем персонажей из файла
// 	if err := loadCharacters(); err != nil {
// 		fmt.Println("Ошибка при загрузке персонажей:", err)
// 		return
// 	}

// 	http.HandleFunc("/", homeHandler)
// 	http.HandleFunc("/register", registerHandler) // Обработчик для регистрации
// 	// Запуск сервера на порту 8080
// 	fmt.Println("Сервер запущен на http://localhost:8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		fmt.Println("Ошибка при запуске сервера:", err)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// Структура для персонажа
type Character struct {
	Name        string // Имя героя
	Description string // Краткое описание героя
	ImageURL    string // Фото персонажа
	ReleaseDate string // Дата выхода героя
	Role        string // Роль в игре
	Class       string // Тип урона персонажа
	Price       int    // Цена персонажа в игре
}

// Структура для пользователя
type User struct {
	Username string // Имя пользователя
	Password string // Пароль пользователя
	Email    string
	ImageURL string // URL изображения профиля
}

// Список персонажей
var characters []Character

// Функция для загрузки персонажей из файла
func loadCharacters() error {
	file, err := os.Open("characters.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&characters); err != nil {
		return err
	}
	return nil
}

var users []User

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Заголовок
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Получаем значение параметра поиска
	searchQuery := r.URL.Query().Get("search")
	var filteredCharacters []Character

	// Фильтруем персонажей по запросу
	for _, character := range characters {
		if searchQuery == "" || strings.Contains(strings.ToLower(character.Name), strings.ToLower(searchQuery)) {
			filteredCharacters = append(filteredCharacters, character)
		}
	}

	// Создание структуры для передачи данных в шаблон
	data := struct {
		Characters  []Character
		SearchQuery string
	}{
		Characters:  filteredCharacters,
		SearchQuery: searchQuery,
	}

	// Загружаем HTML-шаблон из файла
	tmpl, err := template.ParseFiles("html_go.html")
	if err != nil {
		http.Error(w, "Ошибка при загрузке шаблона", http.StatusInternalServerError)
		return
	}

	// Выполнение шаблона и передача данных
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получаем данные из формы
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email") // Получаем электронную почту

		// Проверяем, существует ли уже пользователь с таким именем
		for _, user := range users {
			if user.Username == username {
				http.Error(w, "Пользователь с таким именем уже существует", http.StatusConflict)
				return
			}
		}

		// Добавляем нового пользователя в список
		users = append(users, User{Username: username, Password: password, Email: email, ImageURL: "https://i.pinimg.com/originals/89/c5/a5/89c5a59771b423c2b6bb08cf7c4d17cb.jpg"})

		// Перенаправляем на главную страницу после успешной регистрации
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Обработка метода GET для отображения страницы регистрации
	if r.Method == http.MethodGet {
		// Загружаем HTML-шаблон для регистрации
		tmpl, err := template.ParseFiles("register.html")
		if err != nil {
			http.Error(w, "Ошибка при загрузке шаблона", http.StatusInternalServerError)
			return
		}

		// Выполнение шаблона и передача данных (если нужно)
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Не разрешен", http.StatusMethodNotAllowed)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	if len(users) == 0 {
		// Если нет зарегистрированных пользователей, перенаправляем на страницу регистрации
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Получаем последнего зарегистрировавшегося пользователя
	lastUser := users[len(users)-1]

	// Создание структуры для передачи данных в шаблон
	data := struct {
		User User
	}{
		User: lastUser,
	}

	// Загружаем HTML-шаблон для профиля
	tmpl, err := template.ParseFiles("profile.html")
	if err != nil {
		http.Error(w, "Ошибка при загрузке шаблона", http.StatusInternalServerError)
		return
	}

	// Выполнение шаблона и передача данных
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
	}
}

func main() {
	// Загружаем персонажей из файла
	if err := loadCharacters(); err != nil {
		fmt.Println("Ошибка при загрузке персонажей:", err)
		return
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", registerHandler) // Обработчик для регистрации
	http.HandleFunc("/profile", profileHandler)   // Обработчик для профиля
	// Запуск сервера на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
