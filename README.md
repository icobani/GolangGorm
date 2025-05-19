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




Additional Roles Reimbursement Scheme anlamına gelir. Bu, Birleşik Krallık Ulusal Sağlık Hizmeti (NHS) kapsamında Primary Care Networks (PCNs) tarafından kullanılan bir finansman ve istihdam programıdır.

## ARRS Nedir?

### Additional Roles Reimbursement Scheme (ARRS):
•	2019 yılında başlatılmıştır.
•	İngiltere’deki Primary Care Network (PCN) adı verilen, birkaç GP kliniğinden (aile hekimi grubu) oluşan yapıların, çeşitli yeni sağlık çalışanı rollerini işe almasını teşvik eder.
•	NHS tarafından finanse edilir; PCN’ler bu çalışanlar için maaş geri ödemesi alır.

### Hangi Rolleri Kapsar?

### ARRS, aşağıdaki gibi klinik ve klinik dışı birçok rolü kapsar:
•	Clinical Pharmacist (Klinik Eczacı)
•	Physician Associate (Hekim Yardımcısı)
•	Social Prescribing Link Worker (Sosyal Reçeteleme Görevlisi)
•	First Contact Physiotherapist
•	Care Coordinator
•	Health and Wellbeing Coach
•	Mental Health Practitioner
•	Paramedic
•	Nursing Associate
•	Digital and Transformation Lead

### Amacı Nedir?
•	GP’lerin üzerindeki yükü azaltmak
•	Daha kapsamlı birinci basamak sağlık hizmeti sunmak
•	Hasta bakımını iyileştirmek
•	Uygun uzmanla ilk teması sağlamak (örneğin bel ağrısı olan biri doğrudan fizyoterapiste yönlendirilebilir)

### Özetle:

ARRS, NHS’in geniş bir sağlık çalışanı kadrosunu aile hekimliği sistemine dahil ederek daha etkin, çok disiplinli ve sürdürülebilir bir birinci basamak sağlık hizmeti hedefinin parçasıdır.


