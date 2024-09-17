package cronjob

import (
	"PayWatcher/database"
	"PayWatcher/model"
	"fmt"
	"os"
	"time"

	"github.com/go-co-op/gocron/v2"
	_ "github.com/joho/godotenv/autoload"
)

func InitCronJobs() {
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println(err)
	}

	j, err := s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(07, 30, 00),
			),
		),
		gocron.NewTask(
			SendDailyAlert,
		),
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Job ID:", j.ID())

	s.Start()

	select {
	case <-time.After(time.Minute):
	}

	err = s.Shutdown()
	if err != nil {
		fmt.Println(err)
	}
}

func SendDailyAlert() {
	var mail, name string
	var net_amount, gross_amount float64
	var count int64
	db := database.DB
	var MailSender = model.MailSender{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     os.Getenv("EMAIL_PORT"),
		Username: os.Getenv("EMAIL_USERNAME"),
		Password: os.Getenv("EMAIL_PASSWORD"),
	}

	var listPayments = make(map[string]string)

	// Tiene que enviar alertas de los pagos que se van a hacer mañana
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	rows, err := db.Table("users").Select("users.email, payments.name, payments.net_amount, payments.gross_amount").Joins("left join payments on payments.user_id = users.id").Where("DATE(payments.charge_date) = ?", tomorrow).Count(&count).Rows()

	if err != nil {
		fmt.Println(err)
	}

	bodyMsg := "Estos son los pagos que se efectuaran mañana. \r\n"

	if count > 0 {
		for rows.Next() {
			rows.Scan(&mail, &name, &net_amount, &gross_amount)
			paymentDetail := fmt.Sprintf("Pago: %s \r\nPrecio Neto: %.2f \r\nPrecio: %.2f", name, net_amount, gross_amount)
			listPayments[mail] += paymentDetail
		}

		for key, value := range listPayments {
			MailSender.SendMail([]string{key}, "Alerta de pagos", bodyMsg+value)
		}
	} else {
		fmt.Println("No hay pagos para enviar la alerta.")
	}
}
