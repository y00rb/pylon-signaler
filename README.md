# pylon-signaler
Pylon meeting 中使用的Signaler server

## Usage

Reqiure go 1.12 or above
1. Clone code
```
git clone https://github.com/y00rb/pylon-signaler
```
2. Install dependency
```
go get
```
3. Running signaler server in local
```
go run main.go
```
4. Check running status, open `http://localhost:8383/health` browser. if
   succeed, you would see:
```json
{ alive: true}
```
