<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_zh.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_ua.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_ko.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/kr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_pt-br.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_by.md">
        <img height="20px" src="https://raw.githubusercontent.com/Shpota/goxygen/main/.github/flag-by.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_fr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_es.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/es.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_jp.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/jp.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_id.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/id.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_he.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/il.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_tr.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/tr.svg">
    </a>
    <br>
    Goxygen
    <a href="https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild">
        <img src="https://github.com/Shpota/goxygen/workflows/build/badge.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/releases">
        <img src="https://img.shields.io/github/v/tag/shpota/goxygen?color=green&label=version">
    </a>
    <a href="https://gitter.im/goxygen/community">
        <img src="https://badges.gitter.im/goxygen/community.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/pulls">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg">
    </a>
</h1>

<img src="../templates/vue.webapp/src/assets/logo.svg" align="right" width="230px" alt="goxygen logo">

**Go ve Angular, React veya Vue ile bir web projesi olu??turun.**

Goxygen, yeni bir proje olu??tururken sizin zamandan tasarruf etmenize odaklan??r. Sizin i??in b??t??n konfig??rasyonlar?? yap??lm???? bir uygulama iskeleti olu??turur. ???? mant??????n??z?? hemen uygulamaya ba??layabilirsiniz. Goxygen, back-end Go kodu olu??turur, front-end taraf??nda componentlere ba??lar, uygulama i??in bir Dockerfile sa??lar ve production ortam??nda rahat ??al????ma i??in docker-compose dosyalar??n?? olu??turur.

<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>Desteklenen Teknolojiler</b></td>
    </tr>
    </thead>
    <tbody>
    <tr align="center">
        <td align="center">Front End</td>
        <td>Angular</td>
        <td>React</td>
        <td>Vue</td>
    </tr>
    <tr align="center">
        <td>Back End</td>
        <td colspan=3>Go</td>
    </tr>
    <tr align="center">
        <td>Database</td>
        <td>MongoDB</td>
        <td>MySQL</td>
        <td>PostgreSQL</td>
    </tr>
    </tbody>
</table>

## Gereksinimler

Makinenizde Go 1.11 veya daha yenisine sahip olman??z gerekmektedir.

`GO111MODULE` ortam de??i??keni ??retim mant??????n??n ??al????mas?? i??in `auto` olarak ayarlanmal??d??r. Bu ayar, Go 1.15'e kadar olan s??r??mlerde varsay??lan ayard??r. Ancak, Go 1.16 i??in, bunu a????k??a ayarlaman??z gerekmektedir:

```
export GO111MODULE=auto
```

## Nas??l Kullan??l??r

??al????t??r:

```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app

```

`my-app` klas??r??nde bir proje olu??turur.

Varsay??lan olarak React ve MongoDB'yi kullanacakt??r. Farkl?? bir front-end framework'?? ve veritaban??n?? `--frontend` ve `--db` flag'lerini kullanarak se??ebilirsin. ??rne??in bu komut Vue ve PostgreSQL ile bir proje olu??turacakt??r:

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

`--frontend` flag'i `angular`, `react` ve `vue` de??erlerini kabul eder.
`--db` flag'i `mongo`, `mysql` ve `postgres` de??erlerini kabul eder.

Olu??turulan proje `docker-compose` komutu ile ??al????t??rmaya haz??r olacakt??r:

```sh
cd my-app
docker-compose up
```

Build tamamland??ktan sonra, uygulama http://localhost:8080 adresinden eri??ilebilir olacakt??r.

Olu??turulan projeyle nas??l ??al??????laca???? hakk??nda daha fazla ayr??nt??y?? README dosyas??nda bulabilirsiniz.

![Showcase](showcase.gif)

## Olu??turulmu?? bir projenin yap??s?? (??rnek: React/MongoDB)

    my-app
    ????????? server                   # Go proje dosyalar??
    ???   ????????? db                   # MongoDB ileti??imleri
    ???   ????????? model                # domain nesneleri
    ???   ????????? web                  # REST APIs, web server
    ???   ????????? server.go            # Sunucunun ba??lama noktas??
    ???   ????????? go.mod               # Sunucu ba????ml??l??klar??
    ????????? webapp
    ???   ????????? public               # ikonlar, sabit dosyalar, ve index.html
    ???   ????????? src
    ???   ???   ????????? App.js           # Ana React komponenti
    ???   ???   ????????? App.css          # Uygulama bile??enine ??zg?? stiller
    ???   ???   ????????? index.js         # Uygulaman??n giri?? noktas??
    ???   ???   ????????? index.css        # global stiller
    ???   ????????? package.json         # front end ba????ml??l??klar??
    ???   ????????? .env.development     # Geli??tirme ortam?? i??in API Endpointlerini bar??nd??r??r.
    ???   ????????? .env.production      # Production ortam?? i??in API endpoint'leri
    ????????? Dockerfile               # Back-end ve front-end'i birlikte olu??turur.
    ????????? docker-compose.yml       # ??retim ortam?? da????t??m tan??mlay??c??s??
    ????????? docker-compose-dev.yml   # Geli??tirme ihtiya??lar?? i??in yerel MongoDB ??al????t??r??r.
    ????????? init-db.js               # Test verileriyle bir MongoDB koleksiyonu olu??turur.
    ????????? .dockerignore            # Docker yap??lar??nda yok say??lan dosyalar?? belirtir.
    ????????? .gitignore
    ????????? README.md                # Olu??turulan repo'nun nas??l kullan??laca????na dair k??lavuz.

Birim testleri veya ??rnek komponentler gibi dosyalar, basitlik a????s??ndan buraya dahil edilmemi??tir.

## Ba????ml??l??klar

Goxygen, bir projenin temel yap??s??n?? olu??turur ve sizi belirli bir tak??m ara??lar?? kullanmaya zorlamaz. Bu y??zden projenize gereksiz ba????ml??l??klar getirmez. React ve Vue projelerinde yaln??zca arka u?? taraf??nda bir veritaban?? s??r??c??s?? ve [axios] (https://github.com/axios/axios) kullan??r. Angular projeleri yaln??zca Angular'a ??zg?? k??t??phaneleri kullan??r.

## Nas??l katk??da bulunulur

Bir hata bulduysan??z veya projeyi nas??l iyile??tirece??iniz konusunda bir fikriniz varsa [issue olu??turun](https://github.com/Shpota/goxygen/issues) ve m??mk??n olan en k??sa s??re i??erisinde d??zeltece??iz. Ayr??ca Pull Request ile ??nerebilirsin. Repository'i fork'lay??n, de??i??ikleri yap??n ve bize bir Pull Request g??nderin ve k??sa s??re i??erisinde inceleyece??iz. Ayr??ca t??m de??i??ikleri tart????t??????m??z [Gitter chat](https://gitter.im/goxygen/community) var.

## Jenerik

Goxygen'in logosu [Egon Elbre](https://twitter.com/egonelbre) taraf??ndan yarat??lm????t??r.
