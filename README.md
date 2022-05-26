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
