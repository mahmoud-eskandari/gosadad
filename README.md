# Sadad golang gateway
## درگاه پرداخت گولنگ سداد (بانک ملی)

# Usage

## Installation
`go get github.com/mahmoud-eskandari/gosadad`
## Import
```
import(
    "github.com/mahmoud-eskandari/gosadad"
)
```

## Get a token
```
//LocalDateTime,SignData is optional
Token,err := gosadad.GetToken(gosadad.TokenRequest{},"{KEY}")

res,err := gosadad.Verify(Token,"{KEY}")

```
