# Orders Service A

This service handles CSV file, reads its data and then sends the data to the service [orders-b](https://github.com/MalukiMuthusi/orders-b)

The service determines the order's country based on the following REGEX

| Country    | Code | REGEX                    |
| ---------- | ---- | ------------------------ |
| Cameroon   | +237 | \(237\)\ ?[2368]\d{7,8}$ |
| Ethiopia   | +251 | \(251\)\ ?[1-59]\d{8}$   |
| Morocco    | +212 | \(212\)\ ?[5-9]\d{8}$    |
| Mozambique | +258 | \(258\)\ ?[28]\d{7,8}$   |
| Uganda     | +256 | \(256\)\ ?\d{9}$         |

## Country codes

The mapping of the phone number regexp to the country is passed to the application through a csv that looks as follows

```csv
Cameroon, `(237)\ ?[2368]\d{7,8}$`
Ethiopia, `(251)\ ?[1-59]\d{8}$`
Morocco, `(212)\ ?[5-9]\d{8}$`
Mozambique, `(258)\ ?[28]\d{7,8}$`
Uganda, `(256)\ ?\d{9}$`
```

The application uses that data to determine the country of origin for an order based on its phone number

## References

[1]: https://pkg.go.dev/encoding/csv
[2]: https://github.com/google/re2/wiki/Syntax
