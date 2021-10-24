# Assignment
Create a microservice to return a signal based on the Simple Moving Average (SMA) indicator for any given exchange, market and period.
In this example we will work with Coinbase Pro and BTC-EUR market, but we should be options to send to the service.

Binance Chart: https://www.binance.com/en/trade/BTC_EUR?layout=pro
Valid periods: 1m, 5m, 15m, 30m, 1h, 2h, 4h, 1d

## Get Charts
URL: http://cryptohopper-ticker-frontend.us-east-1.elasticbeanstalk.com/v1/{exchange}/candles?pair={market}&start={start_unix_timestamp}&end={end_unix_timestamp}&period={period}
Example: http://cryptohopper-ticker-frontend.us-east-1.elasticbeanstalk.com/v1/coinbasepro/candles?pair=BTC-EUR&start=1621371923&end=1621425923&period=30m

## SMA indicator
The simple moving average indicator uses the "close" values of the chart.
For the signal we need to use the SMA with a length of 8 and a length of 55. That means that for the first SMA we will use the last 8 close values for our average, and 55, 55 close values.

Docs: https://www.investopedia.com/terms/s/sma.asp#:~:text=A%20simple%20moving%20average%20(SMA)%20is%20an%20arithmetic%20moving%20average,periods%20in%20the%20calculation%20average.

## Signal calculation
If the SMA(8) is currently lower than SMA(55), but previous value was higher or the same, signal a sell.
If the SMA(8) is higher than SMA(55) and previous value was lower or same, signal a buy.
Otherwise signal neutral.

## Get Ticker Prices
For a more accurate live signal, the last close value needs to be replaced with tha "last" value from the ticker.
URL: http://cryptohopper-ticker-frontend.us-east-1.elasticbeanstalk.com/v1/{exchange}/ticker
Example: http://cryptohopper-ticker-frontend.us-east-1.elasticbeanstalk.com/v1/coinbasepro/ticker


---

#Application walk-through
## Assumptions made
* The implementation for each exchange can be different (eg: different url, different data format).
* The application can support different algorithms/logic/indicators to calculate the signal.

##Folder structure
    .
    └──  app                   # application folder
      ├── config               # server config and routes 
      ├── controllers          # end-points handlers
      ├── exchanges            # contains the implementation for all the supported exchanges
      |  └── coinbase          # coinbase implementation
      ├── indicators           # contains the implementation for all the indicators supported
      |  └── sma               # sma implementation
      └── server               # http server initialization
##Exchanges

The current implementation supports only coinbase, since that was the only url provided, but this can be easily expanded.

### Adding a new Exchange
To add a new exchange:

- create a new folder inside the folder "exchanges" with the name of the exchange
- create a new file in the new exchange folder
- create a new struct that implements the exchange interface
- register the new exchange in the supportedExchanges variable in signal.go, the key will be used to match incoming requests for this exchange

## Get it running

* Create a new directory and clone the project
     ```
    mkdir {folder/path} &&
    cd {folder/path} &&
    git clone https://github.com/Riccardo-merli94/Cryptohopper-assignment.git . 
    ```

### - Docker
You can run the following command to build and run the container. "hopper" is the container name,
you can replace it with anything else. NB: the script will remove any container with the same name,
so make sure the name is not used and that there is no container running on port 8080
```
chmod +x run && ./run hopper
```
then you can request

http://localhost:8080/api/v1/{exchange}/?market={market}&period={period}

http://localhost:8080/api/v1/coinbasepro/?market=BTC-EUR&period=5m

### - Without Docker
Alternatively you can run it directly in you machine if you have go already installed.
```
cd app && go run main.go
```
or
```
cd app && go build -o build/hopper .
./build/hopper
```
