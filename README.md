# colimit

<a href="https://www.buymeacoffee.com/virtcanhead" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>
[![Build Status](https://travis-ci.org/virtcanhead/colimit.svg?branch=master)](https://travis-ci.org/virtcanhead/colimit)

concurrency limiter for Echo framework

## Usage

### Integration

```go
func main() {
  // ...

  e := echo.New()
  e.Use(colimit.New(512))

  // ...
}
```

## License

canhead <hi@canhead.xyz> MIT License
