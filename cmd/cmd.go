package cmd

import (
	"PayWatcher/database"
	"PayWatcher/model"
	"fmt"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func SendAlert() {
	var mail, name string
	var net_amount, gross_amount float64
	db := database.DB
	var MailSender = model.MailSender{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     os.Getenv("EMAIL_PORT"),
		Username: os.Getenv("EMAIL_USERNAME"),
		Password: os.Getenv("EMAIL_PASSWORD"),
	}

	var listPayments = make(map[string]string)

	// Tiene que enviar alertas de los pagos que se van a hacer mañana
	tomorrow := time.Now().Add(24 * time.Hour)

	rows, err := db.Table("users").Select("users.email, payments.name, payments.net_amount, payments.gross_amount").Joins("left join payments on payments.user_id = users.id").Where("payments.charge_date = ?", tomorrow).Rows()

	if err != nil {
		fmt.Println(err)
	}

	bodyMsg := "Estos son los pagos que se efectuaran mañana. \r\n"

	for rows.Next() {
		rows.Scan(&mail, &name, &net_amount, &gross_amount)
		paymentDetail := fmt.Sprintf("Pago: %s \r\nPrecio Neto: %.2f \r\nPrecio: %.2f", name, net_amount, gross_amount)
		listPayments[mail] += paymentDetail
	}

	for key, value := range listPayments {
		MailSender.SendMail([]string{key}, "Alerta de pagos", bodyMsg+value)
	}

}
