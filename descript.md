# BlockChain

#### Block

블록체인을 구성하는 각각의 블록을 의미한다.

블록은 블록헤더, 거래정보, 기타정보로 구성된다.


#### 블록 헤더

6가지의 바이트 형식의 직렬화된 데이터를 포함한다.

- version : 블록 버전 번호를 나타낸다. GolockChain에는 구현되어있지 않다.
- merkle hash : 블록의 유일성을 표시하는 SHA256으로 해싱된 데이터를 말한다. GolockChain의 `Hash` 가 이 역할을 한다.
- previous blockhash : 이전 블록의 merkle hash를 표시한다. GolockChain의 `PrevBlockHash` 가 이 역할을 한다.
- time : 채굴자가 블록을 생성한 시간을 표시한다. 이 시간 데이터는 이전 블록보다 반드시 커야한다. GolockChain의 `Timestamp` 가 이 역할을 한다.
- bits : 블록을 생성하는데 요구되는 생성 조건을 표시한다. 좀 더 쉽게 설명하면 생성되는 블록의 merkle hash데이터는 반드시 bits 데이터의 크기 보다 작아야하는데, 랜덤으로 생성된 merkle hash 데이터가 bits보다 클 경우 해당 merkle hash값은 무효화 되며 다시 계산되어야한다. GolockChain에서 `Target` 이 이 역할을 한다.
- nonce : merkls hash가 유효하다고 판단될 때까지 반복된 계산 횟수를 말한다. GolockChain에서 `nonce`가 이 역할을 한다.