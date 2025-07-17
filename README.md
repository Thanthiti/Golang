# Golang
Pratice all basic golang for Cooperative Education 


Array ไม่ได้ส่งแบบ Pass by Reference แต่เป็น Slice แทนที่สามารถใช้หลักการนี้ได่

ถ้าจะแปลงเป็น Json ต้องเป็นพิมพ์ใหญ่ เช่น First ไม่ใช่ first , Last (last)

---------------- Import Package for start project ----------------
// create file go.mod
go mod init (nameModule)                            

// go framework
go get github.com/gin-gonic/gin

// log
go get github.com/sirupsen/logrus                   

// Live server 
go install github.com/air-verse/air@latest          
---------------- Posgres ---------------- 
go get github.com/lib/pq

go get gorm.io/driver/sqlite 


---------------- Fiber ----------------
go get "github.com/gofiber/fiber/v2"