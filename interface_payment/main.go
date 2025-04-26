package main

import (
	"errors"
	"fmt"
	"strings"
)

// Interface: PaymentProcessor
type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
	RefundPayment(amount float64) bool
}

// Abstract-like structure: OnlinePaymentProcessor
type OnlinePaymentProcessor struct {
	ApiKey string
}

func (o *OnlinePaymentProcessor) ProcessPayment(amount float64) bool {
	if !o.ValidateApiKey() {
		panic(errors.New("Invalid API Key"))
	}
	return o.ExecutePayment(amount)
}

func (o *OnlinePaymentProcessor) RefundPayment(amount float64) bool {
	if !o.ValidateApiKey() {
		panic(errors.New("Invalid API Key"))
	}
	return o.ExecuteRefund(amount)
}

// These methods must be implemented by child structs
func (o *OnlinePaymentProcessor) ValidateApiKey() bool {
	panic("ValidateApiKey() not implemented")
}
func (o *OnlinePaymentProcessor) ExecutePayment(amount float64) bool {
	panic("ExecutePayment() not implemented")
}
func (o *OnlinePaymentProcessor) ExecuteRefund(amount float64) bool {
	panic("ExecuteRefund() not implemented")
}

// StripeProcessor
type StripeProcessor struct {
	OnlinePaymentProcessor
}

func (s *StripeProcessor) ProcessPayment(amount float64) bool {
	if !s.ValidateApiKey() {
		panic(errors.New("invalid API Key"))
	}
	return s.ExecutePayment(amount)
}

func (s *StripeProcessor) RefundPayment(amount float64) bool {
	if !s.ValidateApiKey() {
		panic(errors.New("invalid API Key"))
	}
	return s.ExecuteRefund(amount)
}

func (s *StripeProcessor) ValidateApiKey() bool {
	return strings.HasPrefix(s.ApiKey, "sk_")
}

func (s *StripeProcessor) ExecutePayment(amount float64) bool {
	fmt.Printf("Processing Stripe payment of $%.2f\n", amount)
	return true
}

func (s *StripeProcessor) ExecuteRefund(amount float64) bool {
	fmt.Printf("Processing Stripe refund of $%.2f\n", amount)
	return true
}

type PaypalProcessor struct {
	OnlinePaymentProcessor
}

func (p *PaypalProcessor) ProcessPayment(amount float64) bool {
	if !p.ValidateApiKey() {
		panic(errors.New("invalid API Key"))
	}
	return p.ExecutePayment(amount)
}

func (p *PaypalProcessor) RefundPayment(amount float64) bool {
	if !p.ValidateApiKey() {
		panic(errors.New("invalid API Key"))
	}
	return p.ExecuteRefund(amount)
}

func (p *PaypalProcessor) ValidateApiKey() bool {
	return len(p.ApiKey) == 32
}

func (p *PaypalProcessor) ExecutePayment(amount float64) bool {
	fmt.Printf("Processing Paypal payment of $%.2f\n", amount)
	return true
}

func (p *PaypalProcessor) ExecuteRefund(amount float64) bool {
	fmt.Printf("Processing Paypal refund of $%.2f\n", amount)
	return true
}

// CashPaymentProcessor (directly implements PaymentProcessor)
type CashPaymentProcessor struct{}

func (c *CashPaymentProcessor) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing Cash Payment $%.2f\n", amount)
	return true
}

func (c *CashPaymentProcessor) RefundPayment(amount float64) bool {
	fmt.Printf("Processing Cash refund of $%.2f\n", amount)
	return true
}

// OrderProcessor
type OrderProcessor struct {
	PaymentProcessor PaymentProcessor
}

func (o *OrderProcessor) ProcessOrder(amount float64) {
	if o.PaymentProcessor.ProcessPayment(amount) {
		fmt.Println("Order processed successfully")
	} else {
		fmt.Println("Order processing failed")
	}
}

func (o *OrderProcessor) RefundOrder(amount float64) {
	if o.PaymentProcessor.RefundPayment(amount) {
		fmt.Println("Order refunded successfully")
	} else {
		fmt.Println("Order refund failed")
	}
}

func main() {
	stripeProcessor := &StripeProcessor{
		OnlinePaymentProcessor{ApiKey: "sk_test_123456"},
	}
	paypalProcessor := &PaypalProcessor{
		OnlinePaymentProcessor{ApiKey: "valid_paypal_api_key_32charslong"},
	}
	cashProcessor := &CashPaymentProcessor{}

	stripeOrder := &OrderProcessor{PaymentProcessor: stripeProcessor}
	paypalOrder := &OrderProcessor{PaymentProcessor: paypalProcessor}
	cashOrder := &OrderProcessor{PaymentProcessor: cashProcessor}

	stripeOrder.ProcessOrder(100.00)
	paypalOrder.ProcessOrder(150.00)
	cashOrder.ProcessOrder(50.00)

	stripeOrder.RefundOrder(25.00)
	paypalOrder.RefundOrder(50.00)
	cashOrder.RefundOrder(10.00)
}
