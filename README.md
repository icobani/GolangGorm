# GolangGorm

![Logo](/assets/Logo.jpg)


# Projenin açılması

``` go
go mod init github.com/icobani/GolangGorm
```


ile başlar ve yazmış olduğunuz modülün / paketin ismini vermelisiniz.
Bu size tekrar kullanılabilirlik anlamında ulaşılabilir fonksiyonlar yazmanıza olanak sağlar.


kendi makinanıza postgresql kurabileceğiniz gibi hızlıca bir docker container ile de bunu sağlayabilirsiniz.



``` shell
docker run --name my-postgres \
  -e POSTGRES_PASSWORD=secret123 \
  -e POSTGRES_USER=myuser \
  -e POSTGRES_DB=mydb \
  -p 5432:5432 \
  -d postgres:15
```

## Projenin Amacı :

Ingiltere Sağlık sisteminde devlet teşviklerinin hastanelere dağıtılması için personel bazında bildirimlerin yapılması ve takip edilmesi gerekmektedir. Bunları kullanıcı giriş yaparak kendi hastaneleri için bu tanımların yapılabilmesi sağlanacaktır.