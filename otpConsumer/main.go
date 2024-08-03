package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	otpservice "otpServiceConsumer/services/otpService"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)
var OTPServiceMessageBrokerValues=map[string]string{
    "exchange":"otpServiceExchange",
    "queue":"otpServiceQueue",
}

type ConsumerMessageType struct {
	PhoneNumber string `validate:"required,len=10"`
	UserId int `validate:"required"`
}
func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}
// keep listening here
func main(){
	godotenv.Load()
	conn, err := amqp.Dial(os.Getenv("AMQP_CLOUD_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")
	queue,err:=ch.QueueDeclare(OTPServiceMessageBrokerValues["queue"],true,false,false,false,nil)
	fmt.Println(queue)
	failOnError(err,fmt.Sprintf("Failed to declare queue %s",OTPServiceMessageBrokerValues["queue"]))
	
    err=ch.ExchangeDeclare(OTPServiceMessageBrokerValues["exchange"],"fanout",true,false,false,false,nil)
	failOnError(err,"Failed to declare exchange")



failOnError(err, "Failed to open a channel")
defer ch.Close()

failOnError(err, "Failed to declare a queue")
msgs, err := ch.Consume(
  OTPServiceMessageBrokerValues["queue"], // queue
  "",     // consumer
  true,   // auto-ack
  false,  // exclusive
  false,  // no-local
  false,  // no-wait
  nil,    // args
)
failOnError(err, "Failed to register a consumer")

var forever chan struct{}

otpService:=otpservice.InitializeOTPService(os.Getenv("DSN"))
go func() {
  for d := range msgs {
	// parse a message
	println(msgs)
	var consumerMessage ConsumerMessageType;

err:=json.Unmarshal(d.Body,&consumerMessage)
failOnError(err,"Error parsing consumer message")
fmt.Printf("consumer message is %v",consumerMessage.UserId)
otpService.SendOtpWith2Factor(os.Getenv("2FA_AUTHENTICATION_URI"),consumerMessage.UserId,consumerMessage.PhoneNumber)
    log.Printf("Received a message: %s", d.Body)
  }
}()

log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
<-forever
}