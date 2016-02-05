<h1 align="center">Domainchecker</h1>

<p align="center">
Domain checking from the terminal at your favorite supplier. Just opens your default browser when choosed.
</p>

<p align="center">
    <a href="http://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/npm/l/express.svg">
    </a>    
    <a href="https://github.com/KeizerDev/domainchecker/releases/tag/v0.0.1">
        <img src="http://img.shields.io/badge/release-v0.0.1-1eb0fc.svg">
    </a>
</p>

----

```
Usage:
  s <query> [flags]

Flags:
  -b, --binary string     binary to launch search uri
  -l, --list-providers    list supported providers
  -p, --provider string   set search provider (default "google")
  -v, --verbose           display url when opening
      --version           display version
```

## Install

```
go get -v github.com/KeizerDev/domainchecker
cd $GOPATH/src/github.com/KeizerDev/domainchecker
make
make install
```

## Examples

Try example.com on godaddy.com.
```
domainchecker example.com
```

## Provider Expansion

Just use your preffered supplier using the `-p` tag.
```
domainchecker example.com -p namecheap
```

**Todo:**

You can also change the default provider in your domainchecker config file like this: 
```
default: namecheap.com
```

## Supported Providers

* transip

#### Contributors

* [Robert-Jan Keizer (KeizerDev)](https://github.com/KeizerDev/)

#### License

Domainchecker is released under the MIT license.