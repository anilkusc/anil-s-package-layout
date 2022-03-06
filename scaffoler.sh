#!/bin/bash
# This script stands for scaffolding this package layout from stratch
PACKAGE_NAME=$1
if [[ $PACKAGE_NAME == "" ]]
then
PACKAGE_NAME=sample/sample
fi
mkdir api app domain pkg repository
go mod init github.com/$PACKAGE_NAME
echo '
package main
import (
	"github.com/anilkusc/go-package-layout/app"
)
func main() {
	app := app.App{}
	app.Start()
}
' > main.go
echo '
FROM golang:1.17.1 as build
ENV DB_CONN=sqlite
RUN apt-get update && apt-get install sqlite3 -y
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go test -v -cover ./...
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o /bin/app .
FROM alpine
ENV DB_CONN=sqlite
ENV ENV=local
WORKDIR /app
COPY --from=build /bin/app .
CMD ["./app"]
'> Dockerfile
echo '# Applications'> Readme.md
echo '*test.db*'> .gitignore
echo 'ENV=local'>.env
echo 'package app
type App struct {
}
func (app *App) Init() {
	var err error
	repository := repository.Repository{}
	err = repository.Init()
	if err != nil {panic(err)}
	packg := pkg.Pkg{Repository: &repository,}
	err = packg.Init()
	if err != nil {panic(err)}
	dmn := domain.Domain{Pkg: &packg,}
	err = dmn.Init()
	if err != nil {panic(err)}
	api := api.Api{Domain: &dmn,}
	api.Start()
}
func (app *App) Start() {app.Init()}'> app/app.go
echo 'package api
type Api struct {
	Router       *mux.Router
	SessionStore *sessions.CookieStore
	Domain       *domain.Domain
}
func (api *Api) Init() {
	api.Router = mux.NewRouter()
	api.InitRoutes()
	api.SessionStore = sessions.NewCookieStore([]byte(os.Getenv("STORE_KEY")))
}
func (api *Api) InitRoutes() {
	api.Router.HandleFunc("/user/get", api.GetUserHandler)
}
func (api *Api) Start() {
	api.Init()
	fmt.Println("Serving on: ")
	http.ListenAndServe(":8080", api.Router)
}'>api/api.go
echo 'package api
func Construct() (Api, error) {
	godotenv.Load("../.env")
	api := Api{Domain: &domain.Domain{Pkg: &pkg.Pkg{Repository: &repository.Repository{},}}}
	err := api.Domain.Pkg.Repository.Init()
	if err != nil {return api, err}
	err = api.Domain.Pkg.Init()
	if err != nil {return api, err}
	err = api.Domain.Init()
	if err != nil {return api, err}
	api.Init()
	return api, nil
}
func Destruct(db *repository.Database) {
	db.Sqlite.Exec("DROP TABLE mytable")
}
func TestConstruct(t *testing.T) {
	tests := []struct {err error}{{err: nil},}
	for _, test := range tests {
		api, err := Construct()
		if test.err != err {t.Errorf("Error is: %v . Expected: %v", err, test.err)}
		Destruct(api.Domain.Pkg.Repository.Database)}}
'>api/api_test.go
echo 'package api
func (api *Api) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	result, err := api.Domain.GetUser(r.URL.Query().Get("name"))
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	http.Error(w, strconv.Itoa(result), http.StatusOK)
	return
}'>api/user.go
echo 'package api
func TestGetUserHandler(t *testing.T) {
	api, err := Construct()
	if err != nil {t.Errorf("Error is: %v . Expected: %v", err, nil)}
	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{{input: "anil", output: "anil" + "\n", status: 200, err: nil},}
	for _, test := range tests {
		req, err := http.NewRequest("GET", "/user/get?name="+test.input, strings.NewReader(""))
		if err != nil {	t.Errorf("Error is: %v . Expected: %v", err, test.err)}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(api.GetUserHandler)
		handler.ServeHTTP(rr, req)
		if rr.Result().StatusCode != test.status {t.Errorf("Response status is: %v . Expected: %v", rr.Result().StatusCode, test.status)}
		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != string(test.output) {t.Errorf("Response is: %v . Expected: %v", string(body), test.output)}
	}
	Destruct(api.Domain.Pkg.Repository.Database)
}'>api/user_test.go
echo 'package domain
type Domain struct {
	Pkg *pkg.Pkg
}
func (domain *Domain) Init() error {
	return nil
}'>domain/domain.go
echo 'package domain
func Construct() (Domain, error) {
	godotenv.Load("../.env")
	domain := Domain{Pkg: &pkg.Pkg{Repository: &repository.Repository{},},}
	err := domain.Pkg.Repository.Init()
	if err != nil {	return domain, err }
	err = domain.Pkg.Init()
	if err != nil {	return domain, err	}
	err = domain.Pkg.Repository.Init()
	if err != nil {	return domain, err	}
	return domain, nil
}
func Destruct(db *repository.Database) {
	db.Sqlite.Exec("DROP TABLE users")
}
func TestConstruct(t *testing.T) {
	tests := []struct {err error}{{err: nil},}
	for _, test := range tests {
		domain, err := Construct()
		if test.err != err {t.Errorf("Error is: %v . Expected: %v", err, test.err)}
		Destruct(domain.Pkg.Repository.Database)
	}
}
func TestInit(t *testing.T) {
	d, err := Construct()
	if err != nil {t.Errorf("Error is: %v . Expected: %v", err, nil)}
	tests := []struct {	err error}{	{err: nil},	}
	for _, test := range tests {
		err := d.Init()
		if test.err != err {t.Errorf("Error is: %v . Expected: %v", err, test.err)}
		Destruct(d.Pkg.Repository.Database)
	}
}
'>domain/domain.go
echo 'package domain
func Construct() (Domain, error) {
	godotenv.Load("../.env")
	domain := Domain{Pkg: &pkg.Pkg{Repository: &repository.Repository{},},}
	err := domain.Pkg.Repository.Init()
	if err != nil {	return domain, err	}
	err = domain.Pkg.Init()
	if err != nil {	return domain, err	}
	err = domain.Pkg.Repository.Init()
	if err != nil {	return domain, err }
	return domain, nil
}
func Destruct(db *repository.Database) {
	db.Sqlite.Exec("DROP TABLE factorials")
}
func TestConstruct(t *testing.T) {
	tests := []struct {	err error}{	{err: nil},	}
	for _, test := range tests {
		domain, err := Construct()
		if test.err != err { t.Errorf("Error is: %v . Expected: %v", err, test.err)	}
		Destruct(domain.Pkg.Repository.Database)
	}
}
func TestInit(t *testing.T) {
	d, err := Construct()
	if err != nil {	t.Errorf("Error is: %v . Expected: %v", err, nil)}
	tests := []struct {err error}{{err: nil},}
	for _, test := range tests {
		err := d.Init()
		if test.err != err {t.Errorf("Error is: %v . Expected: %v", err, test.err)}
		Destruct(d.Pkg.Repository.Database)
	}
}'>domain/domain_test.go
echo 'package domain
func (domain *Domain) GetUser(input string) (string, error) {
	user := pkg.User{}
	user.Name = input
	return user.GetUser(domain.Pkg.Repository)
}'>domain/user.go
echo 'package domain
func TestCalculate(t *testing.T) {
	d, err := Construct()
	if err != nil {	t.Errorf("Error is: %v . Expected: %v", err, nil)}
	tests := []struct {
		input  string
		output string
		err    error
	}{{input: "anil", output: "anil", err: nil},}
	for _, test := range tests {
		output, err := d.Calculate(test.input)
		if err != test.err {if err.Error() != test.err.Error() {t.Errorf("Error is: %v . Expected: %v", err, test.err)}}
		if test.output != output {t.Errorf("Result is: %v . Expected: %v", output, test.output)}
	}
	Destruct(d.Pkg.Repository.Database)
}'>domain/user_test.go
echo 'package pkg
type Pkg struct {
	Repository *repository.Repository
}
func (pkg *Pkg) Init() error {
	return pkg.Repository.Database.Sqlite.AutoMigrate(&Factorial{})
}'>pkg/pkg.go
echo 'package pkg
func Construct() (Pkg, Factorial, error) {
	godotenv.Load("../.env")
	packg := Pkg{Repository: &repository.Repository{},}
	user := User{	Model: gorm.Model{UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},},Name:  "anil"}
	err := packg.Repository.Init()
	if err != nil {	return packg, factorial, err}
	err = packg.Init()
	if err != nil {	return packg, factorial, err	}
	return packg, factorial, nil
}
func Destruct(db *repository.Database) {
	db.Sqlite.Exec("DROP TABLE users")
}
func TestConstruct(t *testing.T) {
	tests := []struct {	err error}{	{err: nil},	}
	for _, test := range tests {
		repository, _, err := Construct()
		if test.err != err {t.Errorf("Error is: %v . Expected: %v", err, test.err)}
		Destruct(repository.Repository.Database)
	}
}
func TestInit(t *testing.T) {
	p, _, err := Construct()
	if err != nil {	t.Errorf("Error is: %v . Expected: %v", err, nil) }
	tests := []struct {	err error }{{err: nil},	}
	for _, test := range tests {
		err := p.Init()
		if test.err != err {t.Errorf("Error is: %v . Expected: %v", err, test.err)}
		Destruct(p.Repository.Database)
	}
}'>pkg/pkg_test.go
echo 'package pkg
type User struct {
	gorm.Model
	Name  string
}
func (u *User) GetUser(repository *repository.Repository) (int, error) {
	res := repository.Database.Sqlite.First(u)
	return f.Result, res.Error
}'>pkg/user.go
echo 'package pkg
func TestCalculate(t *testing.T) {
	p, f, err := Construct()
	if err != nil {	t.Errorf("Error is: %v . Expected: %v", err, nil)}
	tests := []struct {
		input  Factorial
		output int
		err    error
	}{{input: f, output: f.Result, err: nil},}
	for _, test := range tests {
		output, err := test.input.Calculate(p.Repository)
		if test.err != err {t.Errorf("Error is: %v . Expected: %v", err, test.err)}
		if test.output != output {t.Errorf("Result is: %v . Expected: %v", output, test.output)}
	}
	Destruct(p.Repository.Database)
}'>pkg/user_test.go
echo 'package repository
type Repository struct {
	Database *Database
}
func (repository *Repository) Init() error {
	var err error
	repository.Database = &Database{}
	repository.Database.Sqlite, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return err
}'>repository/repository.go
echo 'package repository
func Construct() (Repository, interface{}, error) {
	godotenv.Load("../.env")
	repository := Repository{}
	type TestingPurposeStruct struct {
		gorm.Model
		Name string
		Role string
	}
	var tst = TestingPurposeStruct{
		Model: gorm.Model{
			UpdatedAt: time.Time{}, CreatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Name: "test",
		Role: "admin",
	}
	err := repository.Init()
	if err != nil {	return repository, tst, err	}
	err = repository.Database.Sqlite.AutoMigrate(&TestingPurposeStruct{})
	if err != nil {	return repository, tst, nil	}
	return repository, tst, nil
}
func Destruct(db *Database) {
	db.Sqlite.Exec("DROP TABLE testing_purpose_structs")
}
func TestConstruct(t *testing.T) {
	tests := []struct {	err error }{{err: nil},}
	for _, test := range tests {
		repository, _, err := Construct()
		if test.err != err {t.Errorf("Error is: %v . Expected: %v", err, test.err)}
		Destruct(repository.Database)
	}
}
func TestInit(t *testing.T) {
	repository, _, err := Construct()
	if err != nil {	t.Errorf("Error is: %v . Expected: %v", err, nil)	}
	tests := []struct {	err error}{	{err: nil},	}
	for _, test := range tests {
		err := repository.Init()
		if test.err != err {t.Errorf("Error is: %v . Expected: %v", err, test.err)	}
		Destruct(repository.Database)
	}
}
'>repository/repository_test.go
echo 'package repository
type Database struct {
	Sqlite *gorm.DB
}'>repository/database.go