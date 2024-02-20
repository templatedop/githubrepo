package main

import (
	//"context"
	"fmt"
	"time"

	//"time"
	"github.com/templatedop/githubrepo/dtime"
	//"github.com/aarondl/opt/null"
	//"github.com/jackc/pgx/v5"
	//"github.com/jackc/pgx/v5/pgtype"
	//"github.com/jackc/pgx/v5/pgtype"
	//"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	// conn, err := pgx.Connect(context.Background(), "postgresql://postgres:secret@localhost/templatedatabase")
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close(context.Background())

	// // Execute a query
	// rows, err := conn.Query(context.Background(), "SELECT * FROM users")

	// //fmt.Println("Rows:", rows)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// // Define the struct to hold the query result
	// type User struct {
	// 	ID          uint64           `db:"id"`
	// 	Name        null.Val[string] `db:"name"`
	// 	Email       null.Val[string] `db:"email"`
	// 	Password    null.Val[string] `db:"password"`
	// 	CreatedAt   null.Val[string] `db:"created_at"`
	// 	UpdatedAt   null.Val[string] `db:"updated_at"`
	// 	CreatedTime null.Val[string] `db:"created_time"`
	// }

	// type User1 struct {
	// 	ID        uint64           `db:"id"`
	// 	Name      null.Val[string] `db:"name"`
	// 	Email     string           `db:"email"`
	// 	Password  string           `db:"password"`
	// 	CreatedAt time.Time        `db:"created_at"`
	// 	UpdatedAt time.Time        `db:"updated_at"`
	// }

	// //iuser:= User1{Name:"Ram",Password: "sadf",Email: "Ram@yahoo.com",CreatedAt: 2024/02/01}

	// collectedRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	// if err != nil {
	// 	fmt.Println("pgutility, Err collecting Rows at Select Rows :", err.Error())

	// }

	// fmt.Println("Collected Rows:", collectedRows)

	s := "2024-02-25 08:08:08"
	s1 := "08:12"
	dt3 := dtime.NewFromStrLayout(s1, "10:45")
	fmt.Println("dt3:", dt3)

	dt1 := dtime.New(s)
	fmt.Println("Sring time:", dt1)
	dt2, _ := dt1.ToZone("Asia/Kolkata")
	fmt.Println("Indian format: ", dt2.Format("d-m-Y H:i:s"))            //25-02-2024 08:08:08
	fmt.Println("dt2:", dt2)                                             //2024-02-25 08:08:08
	fmt.Println("Month:", dt2.Month())                                   //2
	fmt.Println("Day:", dt2.Day())                                       //25
	fmt.Println("Second:", dt2.Second())                                 //8
	fmt.Println("Time addition", dt2.Add(time.Duration(10)*time.Second)) //2024-02-25 08:08:18
	fmt.Println("time:", dt2.Format("h:i"))                              // 08:08
	fmt.Println("WeekDay:", dt2.Format("l"))                             //Sunday
	fmt.Println("Check equal:", dt1.Equal(dt2))                          //true
	fmt.Println("Start of week:", dt2.StartOfWeek())                     //2024-02-25 00:00:00
	//Similar is the case with end as well.
	fmt.Println("Start of Quarter:", dt2.StartOfQuarter()) //2024-01-01 00:00:00
	sdate, _ := dtime.StrToTime("2024-02-18 14:15:05", "Y-m-d H:i:s")
	fmt.Println("Convert from string to time:", sdate)
	fmt.Println("d D j l:", dt2.Format("d D j l"))                           //25 Sun 25 Sunday
	fmt.Println("F m M n:", dt2.Format("F m M n"))                           //February 02 Feb 2
	fmt.Println("Y y:", dt2.Format("Y y"))                                   //2024 24
	fmt.Println("a A g G h H i s u .u:", dt2.Format("a A g G h H i s u .u")) //am AM 8 8 08 08 08 08 000 .000
	fmt.Println("O P T:", dt2.Format("O P T"))                               //+0530 +05:30 IST
	fmt.Println("r:", dt2.Format("r"))                                       //Sun, 25 Feb 24 08:08 IST
	fmt.Println("c:", dt2.Format("c"))                                       //2024-02-25T08:08:08+05:30
	dtimes := dtime.NewFromStr("2024-02-18 14:15:05")
	fmt.Println("String time: ", dtimes)
	fmt.Println("Format: ", dtime.NewFromStrFormat("18-02-2024 14:15:05", "d-m-Y H:i:s"))
	fmt.Println("Format: ", dtime.NewFromStrFormat("14:15", "H:i"))
	fmt.Println("String format:", dt2.String())
	d, _ := dtime.ParseDuration("14")

	fmt.Println("", d)

}
