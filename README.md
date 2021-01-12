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
Token,err := gosadad.GetToken(sadad.TokenRequest{},"{KEY}")

```