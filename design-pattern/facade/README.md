# Facade pattern
Facade is a structural design pattern that provides a simplified interface to a complex system of classes, library or framework.
## Problem
Design a wallet system with actions:
- Check account
- Check security PIN
- Credit/debit balance
- Make ledger entry
- Send notification
## Solution
In a complex system like this, it’s easy to get lost and easy to break stuff if you’re doing something wrong. That’s why there’s a concept of the Facade pattern: a thing that lets the client work with dozens of components using a simple interface. The client only needs to enter the card details, the security pin, the amount to pay, and the operation type. The Facade directs further communications with various components without exposing the client to internal complexities.
