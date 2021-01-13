# DDD
### 各レイヤー間での依存関係
すべての依存が中心に向かっているこの状態が理想。

![](https://yyh-gl.github.io/tech-blog/img/tech-blog/2019/06/go_web_api/dependency_direction1.png)

基本的に依存はひとつ下の層までに抑える。
プロジェクトの規模やチームの意向によってどこまで依存させるかは、変えても良い場合もある
※正し、下層が上層に依存するのは有り得ない。

### 他のアーキテクチャとの違い
一般的なアーキテクチャだと
ユーザからのリクエストは Handler で受け取られ、UseCaseを使って処理が行われる。さらに、UseCaseはDomainを使って処理を行う。
ここまでは処理が中心に進む。
しかし、DB を使用した場合、UseCaseからDomainを介して、Infra を利用することになる。
UseCase → Domain → Infra (DBを使用したい為に、依存性が外を向く)

![](https://yyh-gl.github.io/tech-blog/img/tech-blog/2019/06/go_web_api/dependency_direction2.png)

DDDではこれを解消する為に、 [依存性逆転の法則](https://medium.com/eureka-engineering/go-dependency-inversion-principle-8ffaf7854a55)  を使う。

### 依存性逆転の仕方
① Domain層において、DBとのやりとりをinterfaceで定義する。
interface自体は実装を持たないので、どこにも依存しない。

② Infra層からDomain層に定義したinterfaceを実装する。

①, ② の2ステップを踏むことで、まずDomain層はinterfaceに対して処理をお願いするだけ良い。 
interfaceは実装を持たないので依存関係はない。

interface自体は実装を持たないが、
Infra層がinterfaceの実装を行っているので、DBにアクセスして処理を行うことができる。
つまりInfra層がinterfaceへの依存となり 、Domain層に向く。

![](https://yyh-gl.github.io/tech-blog/img/tech-blog/2019/06/go_web_api/dependency_direction3.png)

依存性が逆転し、すべての依存関係が中心に向かうようになる。



