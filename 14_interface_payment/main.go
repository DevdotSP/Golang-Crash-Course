package main

import (
	"errors"
	"fmt"
	"strings"
)

// -----------------------------
// INTERFACE DEFINITION
// -----------------------------

// PaymentProcessor defines a contract.
// Any type that implements BOTH ProcessPayment() and RefundPayment()
// is considered a PaymentProcessor (no "implements" keyword needed in Go).
type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
	RefundPayment(amount float64) bool
}

// -----------------------------
// BASE STRUCT (like an abstract class)
// -----------------------------

// OnlinePaymentProcessor acts as a "base" type that Stripe/Paypal embed.
// It provides shared fields (ApiKey) and some default logic,
// but leaves actual implementations (ValidateApiKey, ExecutePayment, ExecuteRefund)
// to child structs.
type OnlinePaymentProcessor struct {
	ApiKey string
}

// These methods "exist" but panic because they must be overridden
// by concrete processors (Stripe, PayPal, etc.)
func (o *OnlinePaymentProcessor) ValidateApiKey() bool {
	panic("ValidateApiKey() not implemented")
}
func (o *OnlinePaymentProcessor) ExecutePayment(amount float64) bool {
	panic("ExecutePayment() not implemented")
}
func (o *OnlinePaymentProcessor) ExecuteRefund(amount float64) bool {
	panic("ExecuteRefund() not implemented")
}

// -----------------------------
// STRIPE PROCESSOR
// -----------------------------

// StripeProcessor "embeds" OnlinePaymentProcessor (composition).
// This means it inherits ApiKey + can override methods.
type StripeProcessor struct {
	OnlinePaymentProcessor
}

// StripeProcessor fulfills the PaymentProcessor interface
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

// Stripe-specific logic
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

// -----------------------------
// PAYPAL PROCESSOR
// -----------------------------

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

// Paypal-specific logic
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

// -----------------------------
// CASH PROCESSOR
// -----------------------------

// CashPaymentProcessor directly implements PaymentProcessor
// without needing the OnlinePaymentProcessor "base type".
type CashPaymentProcessor struct{}

func (c *CashPaymentProcessor) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing Cash Payment $%.2f\n", amount)
	return true
}
func (c *CashPaymentProcessor) RefundPayment(amount float64) bool {
	fmt.Printf("Processing Cash refund of $%.2f\n", amount)
	return true
}

// -----------------------------
// ORDER PROCESSOR
// -----------------------------

// OrderProcessor depends on the PaymentProcessor interface,
// NOT a specific implementation (Stripe, Paypal, Cash).
// This is polymorphism: we can swap processors easily.
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

// -----------------------------
// MAIN FUNCTION
// -----------------------------

func main() {
	// Create different payment processors
	stripeProcessor := &StripeProcessor{
		OnlinePaymentProcessor{ApiKey: "sk_test_123456"},
	}
	paypalProcessor := &PaypalProcessor{
		OnlinePaymentProcessor{ApiKey: "valid_paypal_api_key_32charslong"},
	}
	cashProcessor := &CashPaymentProcessor{}

	// Inject different processors into OrderProcessor
	stripeOrder := &OrderProcessor{PaymentProcessor: stripeProcessor}
	paypalOrder := &OrderProcessor{PaymentProcessor: paypalProcessor}
	cashOrder := &OrderProcessor{PaymentProcessor: cashProcessor}

	// Process payments
	stripeOrder.ProcessOrder(100.00)
	paypalOrder.ProcessOrder(150.00)
	cashOrder.ProcessOrder(50.00)

	// Process refunds
	stripeOrder.RefundOrder(25.00)
	paypalOrder.RefundOrder(50.00)
	cashOrder.RefundOrder(10.00)
}
