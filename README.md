# SLink

Link shortener using firestore

### Link Shorten

```bash
 curl --location --request POST "$URL:8080/api/short" --header 'Content-Type: application/json' --data-raw '{
    "url": "https://www.facebook.com/khanhvu2212/",
    "name": "Custom Name",
    "short": "personal-fb",
    "expired_at": null
}'
```

### Usage
Simply follow the link `$URL/api/g/$SHORT`

Example :`$URL/api/g/personal-fb` will be redirected to `https://www.facebook.com/khanhvu2212/`


### Live Example
[Link](http://www.jsjsjs.ga/)