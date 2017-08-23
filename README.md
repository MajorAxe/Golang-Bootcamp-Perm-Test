# Golang-Bootcamp-Perm-Test

Конвертирует количество (--value) одной валюты (--from) в другую (--to), основываясь на данных сервиса http://fixer.io/

Допустимые валюты:

        "AUD"
        "BGN"
        "BRL"
        "CAD"
        "CHF"
        "CNY"
        "CZK"
        "DKK"
        "GBP"
        "HKD"
        "HRK"
        "HUF"
        "IDR"
        "ILS"
        "INR"
        "JPY"
        "KRW"
        "MXN"
        "MYR"
        "NOK"
        "NZD"
        "PHP"
        "PLN"
        "RON"
        "RUB"
        "SEK"
        "SGD"
        "THB"
        "TRY"
        "USD"
        "ZAR"

Пример использования:

> ./convertCurrency --from="USD" --to="RUB" --value=500
Converted 500 USD to 29514.909523404975 RUB
