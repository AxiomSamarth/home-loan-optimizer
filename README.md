![Build](https://github.com/AxiomSamarth/home-loan-optimizer/actions/workflows/workflow.yaml/badge.svg)
# Table of contents

- [Home Loan Optimizer](#home-loan-optimizer)
- [Requirements and usage](#requirements-and-usage)
  - [Requirements](#requirements)
  - [Usage](#usage)
- [API Design and Documentation](#api-design-and-documentation)
- [Current assumptions in the algorithm](#current-assumptions-in-the-algorithm)
- [Features to be added in future](#features-to-be-added-in-future)
- [Contribution guidelines](#contribution-guidelines)
# Home Loan Optimizer

Owning a home is always an amazing feeling! However, planning the settlement of your home loan in a way that, it is optimally beneficial to you is very important. Home loan does help us save taxes and at the same time it also results in high interests to be paid over a period of time. It is not recommended to prepay the home loan from other funds as one might need that additional funds for something else in life like kid's education, wedding, medical emergency, etc. So, it is not recommended to give away the additional funds that one has to the bank just to reduce interest. 

So, there has to be a way where 
- the borrower avails decent tax benefits
- pays significantly lower interest than the original plan 
- at the same time keep the additional funds with oneself which may be utilised for some other event of life
- also utilize the saved interest to invest in some other instrument to create more wealth

_**This objective can be achieved by simply increasing the EMIs by a certain percentage every year!**_

This project is aimed at recommending you with strategies to achieve the above objective and also provide you with calcuations so that you can arrive at a decision of what suits you the best and increase your EMI by that certain percentage accordingly!

## Requirements and usage

### Requirements
go v1.17

### Usage
Running the API Server

```
# clone the repository
git clone https://github.com/AxiomSamarth/home-loan-optimizer.git

# navigate into the project folder
cd home-loan-optimizer

# start the server which runs by default at http://localhost:8080
go run cmd/main.go
```

Use curl to get home loan optimization recommendations like

```
curl -d '{"loan_amount": 5000000,"annual_interest_rate": 7.15,"loan_tenure": 15}' -H 'Content-Type: application/json' http://localhost:8080/strategies 
```

<details>
	<summary>[toggle] to receive JSON response of type <code>[]StrategyInfo</code> like</summary>

```json
[
  {
    "annual_details": [
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 351253.80645822175,
        "annual_tax_saved": 105376.14193746653,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 1,
            "towards_interest": 29791.666666666668,
            "towards_principle": 15570.093331856096
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 2,
            "towards_interest": 29698.89486056436,
            "towards_principle": 15662.865137958404
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 3,
            "towards_interest": 29605.570289117357,
            "towards_principle": 15756.189709405407
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 4,
            "towards_interest": 29511.68965876549,
            "towards_principle": 15850.070339757276
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 5,
            "towards_interest": 29417.249656324435,
            "towards_principle": 15944.51034219833
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 6,
            "towards_interest": 29322.246948868837,
            "towards_principle": 16039.513049653928
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 7,
            "towards_interest": 29226.678183614647,
            "towards_principle": 16135.081814908117
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 8,
            "towards_interest": 29130.53998780082,
            "towards_principle": 16231.220010721943
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 9,
            "towards_interest": 29033.82896857027,
            "towards_principle": 16327.931029952495
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 10,
            "towards_interest": 28936.541712850132,
            "towards_principle": 16425.218285672632
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 11,
            "towards_interest": 28838.674787231335,
            "towards_principle": 16523.08521129143
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 12,
            "towards_interest": 28740.2247378474,
            "towards_principle": 16621.535260675366
          }
        ],
        "year": 1
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 336986.53054329316,
        "annual_tax_saved": 101095.95916298794,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 13,
            "towards_interest": 28641.18809025254,
            "towards_principle": 16720.571908270224
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 14,
            "towards_interest": 28541.561349299096,
            "towards_principle": 16820.19864922367
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 15,
            "towards_interest": 28441.340999014137,
            "towards_principle": 16920.418999508627
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 16,
            "towards_interest": 28340.523502475397,
            "towards_principle": 17021.236496047368
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 17,
            "towards_interest": 28239.10530168645,
            "towards_principle": 17122.654696836315
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 18,
            "towards_interest": 28137.082817451133,
            "towards_principle": 17224.67718107163
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 19,
            "towards_interest": 28034.45244924725,
            "towards_principle": 17327.307549275516
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 20,
            "towards_interest": 27931.21057509948,
            "towards_principle": 17430.549423423283
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 21,
            "towards_interest": 27827.353551451582,
            "towards_principle": 17534.406447071182
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 22,
            "towards_interest": 27722.877713037782,
            "towards_principle": 17638.882285484982
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 23,
            "towards_interest": 27617.779372753434,
            "towards_principle": 17743.98062576933
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 24,
            "towards_interest": 27512.054821524893,
            "towards_principle": 17849.70517699787
          }
        ],
        "year": 2
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 321665.0415973911,
        "annual_tax_saved": 96499.51247921732,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 25,
            "towards_interest": 27405.700328178613,
            "towards_principle": 17956.05967034415
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 26,
            "towards_interest": 27298.71213930948,
            "towards_principle": 18063.047859213286
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 27,
            "towards_interest": 27191.086479148333,
            "towards_principle": 18170.67351937443
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 28,
            "towards_interest": 27082.81954942873,
            "towards_principle": 18278.940449094036
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 29,
            "towards_interest": 26973.907529252872,
            "towards_principle": 18387.852469269892
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 30,
            "towards_interest": 26864.346574956806,
            "towards_principle": 18497.41342356596
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 31,
            "towards_interest": 26754.132819974722,
            "towards_principle": 18607.627178548042
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 32,
            "towards_interest": 26643.26237470254,
            "towards_principle": 18718.497623820225
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 33,
            "towards_interest": 26531.731326360616,
            "towards_principle": 18830.02867216215
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 34,
            "towards_interest": 26419.535738855648,
            "towards_principle": 18942.224259667117
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 35,
            "towards_interest": 26306.671652641795,
            "towards_principle": 19055.08834588097
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 36,
            "towards_interest": 26193.135084580925,
            "towards_principle": 19168.62491394184
          }
        ],
        "year": 3
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 305211.4435230896,
        "annual_tax_saved": 91563.43305692686,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 37,
            "towards_interest": 26078.922027802022,
            "towards_principle": 19282.837970720742
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 38,
            "towards_interest": 25964.02845155981,
            "towards_principle": 19397.731546962954
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 39,
            "towards_interest": 25848.45030109249,
            "towards_principle": 19513.309697430275
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 40,
            "towards_interest": 25732.183497478632,
            "towards_principle": 19629.576501044132
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 41,
            "towards_interest": 25615.223937493243,
            "towards_principle": 19746.53606102952
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 42,
            "towards_interest": 25497.567493462942,
            "towards_principle": 19864.192505059822
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 43,
            "towards_interest": 25379.210013120293,
            "towards_principle": 19982.54998540247
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 44,
            "towards_interest": 25260.14731945727,
            "towards_principle": 20101.612679065493
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 45,
            "towards_interest": 25140.375210577844,
            "towards_principle": 20221.38478794492
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 46,
            "towards_interest": 25019.889459549668,
            "towards_principle": 20341.870538973097
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 47,
            "towards_interest": 24898.68581425495,
            "towards_principle": 20463.074184267814
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 48,
            "towards_interest": 24776.75999724036,
            "towards_principle": 20585.000001282406
          }
        ],
        "year": 4
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 287542.0844584101,
        "annual_tax_saved": 86262.62533752303,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 49,
            "towards_interest": 24654.107705566046,
            "towards_principle": 20707.652292956718
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 50,
            "towards_interest": 24530.724610653848,
            "towards_principle": 20831.035387868917
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 51,
            "towards_interest": 24406.606358134464,
            "towards_principle": 20955.1536403883
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 52,
            "towards_interest": 24281.748567693812,
            "towards_principle": 21080.011430828952
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 53,
            "towards_interest": 24156.14683291846,
            "towards_principle": 21205.613165604304
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 54,
            "towards_interest": 24029.796721140065,
            "towards_principle": 21331.9632773827
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 55,
            "towards_interest": 23902.693773278996,
            "towards_principle": 21459.06622524377
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 56,
            "towards_interest": 23774.833503686918,
            "towards_principle": 21586.926494835847
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 57,
            "towards_interest": 23646.21139998852,
            "towards_principle": 21715.548598534246
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 58,
            "towards_interest": 23516.822922922252,
            "towards_principle": 21844.937075600512
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 59,
            "towards_interest": 23386.66350618013,
            "towards_principle": 21975.096492342633
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 60,
            "towards_interest": 23255.72855624659,
            "towards_principle": 22106.031442276173
          }
        ],
        "year": 5
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 268567.1314817603,
        "annual_tax_saved": 80570.13944452809,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 61,
            "towards_interest": 23124.013452236362,
            "towards_principle": 22237.746546286402
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 62,
            "towards_interest": 22991.513545731406,
            "towards_principle": 22370.246452791358
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 63,
            "towards_interest": 22858.224160616857,
            "towards_principle": 22503.535837905907
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 64,
            "towards_interest": 22724.140592916003,
            "towards_principle": 22637.61940560676
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 65,
            "towards_interest": 22589.25811062426,
            "towards_principle": 22772.501887898503
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 66,
            "towards_interest": 22453.5719535422,
            "towards_principle": 22908.188044980565
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 67,
            "towards_interest": 22317.07733310752,
            "towards_principle": 23044.682665415243
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 68,
            "towards_interest": 22179.76943222609,
            "towards_principle": 23181.990566296674
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 69,
            "towards_interest": 22041.643405101906,
            "towards_principle": 23320.11659342086
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 70,
            "towards_interest": 21902.69437706611,
            "towards_principle": 23459.065621456655
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 71,
            "towards_interest": 21762.917444404928,
            "towards_principle": 23598.842554117837
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 72,
            "towards_interest": 21622.307674186643,
            "towards_principle": 23739.452324336122
          }
        ],
        "year": 6
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 248190.11389169956,
        "annual_tax_saved": 74457.03416750986,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 73,
            "towards_interest": 21480.860104087475,
            "towards_principle": 23880.89989443529
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 74,
            "towards_interest": 21338.569742216463,
            "towards_principle": 24023.1902563063
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 75,
            "towards_interest": 21195.431566939304,
            "towards_principle": 24166.32843158346
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 76,
            "towards_interest": 21051.440526701117,
            "towards_principle": 24310.319471821647
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 77,
            "towards_interest": 20906.591539848185,
            "towards_principle": 24455.16845867458
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 78,
            "towards_interest": 20760.87949444858,
            "towards_principle": 24600.880504074183
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 79,
            "towards_interest": 20614.299248111805,
            "towards_principle": 24747.46075041096
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 80,
            "towards_interest": 20466.845627807274,
            "towards_principle": 24894.91437071549
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 81,
            "towards_interest": 20318.51342968176,
            "towards_principle": 25043.246568841005
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 82,
            "towards_interest": 20169.29741887575,
            "towards_principle": 25192.462579647014
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 83,
            "towards_interest": 20019.192329338686,
            "towards_principle": 25342.567669184078
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 84,
            "towards_interest": 19868.192863643133,
            "towards_principle": 25493.56713487963
          }
        ],
        "year": 7
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 226307.4327395191,
        "annual_tax_saved": 67892.22982185573,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 85,
            "towards_interest": 19716.293692797808,
            "towards_principle": 25645.466305724956
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 86,
            "towards_interest": 19563.489456059528,
            "towards_principle": 25798.270542463237
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 87,
            "towards_interest": 19409.774760744018,
            "towards_principle": 25951.985237778747
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 88,
            "towards_interest": 19255.144182035587,
            "towards_principle": 26106.615816487178
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 89,
            "towards_interest": 19099.592262795686,
            "towards_principle": 26262.16773572708
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 90,
            "towards_interest": 18943.113513370314,
            "towards_principle": 26418.64648515245
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 91,
            "towards_interest": 18785.70241139628,
            "towards_principle": 26576.057587126485
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 92,
            "towards_interest": 18627.35340160632,
            "towards_principle": 26734.406596916444
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 93,
            "towards_interest": 18468.060895633025,
            "towards_principle": 26893.69910288974
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 94,
            "towards_interest": 18307.81927181164,
            "towards_principle": 27053.940726711124
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 95,
            "towards_interest": 18146.622874981655,
            "towards_principle": 27215.13712354111
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 96,
            "towards_interest": 17984.466016287224,
            "towards_principle": 27377.29398223554
          }
        ],
        "year": 8
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 202807.83412104592,
        "annual_tax_saved": 60842.350236313774,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 97,
            "towards_interest": 17821.3429729764,
            "towards_principle": 27540.417025546365
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 98,
            "towards_interest": 17657.24798819919,
            "towards_principle": 27704.512010323575
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 99,
            "towards_interest": 17492.175270804342,
            "towards_principle": 27869.584727718422
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 100,
            "towards_interest": 17326.11899513502,
            "towards_principle": 28035.641003387744
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 101,
            "towards_interest": 17159.07330082317,
            "towards_principle": 28202.686697699595
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 102,
            "towards_interest": 16991.03229258271,
            "towards_principle": 28370.727705940055
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 103,
            "towards_interest": 16821.990040001485,
            "towards_principle": 28539.76995852128
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 104,
            "towards_interest": 16651.940577331963,
            "towards_principle": 28709.8194211908
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 105,
            "towards_interest": 16480.8779032807,
            "towards_principle": 28880.882095242065
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 106,
            "towards_interest": 16308.79598079655,
            "towards_principle": 29052.96401772621
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 107,
            "towards_interest": 16135.688736857597,
            "towards_principle": 29226.071261665165
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 108,
            "towards_interest": 15961.550062256843,
            "towards_principle": 29400.20993626592
          }
        ],
        "year": 9
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 177571.84354983034,
        "annual_tax_saved": 53271.553064949105,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 109,
            "towards_interest": 15786.373811386591,
            "towards_principle": 29575.386187136173
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 110,
            "towards_interest": 15610.15380202157,
            "towards_principle": 29751.606196501194
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 111,
            "towards_interest": 15432.883815100751,
            "towards_principle": 29928.876183422013
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 112,
            "towards_interest": 15254.557594507864,
            "towards_principle": 30107.2024040149
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 113,
            "towards_interest": 15075.168846850609,
            "towards_principle": 30286.591151672154
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 114,
            "towards_interest": 14894.711241238565,
            "towards_principle": 30467.0487572842
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 115,
            "towards_interest": 14713.178409059743,
            "towards_principle": 30648.58158946302
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 116,
            "towards_interest": 14530.56394375586,
            "towards_principle": 30831.196054766904
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 117,
            "towards_interest": 14346.861400596208,
            "towards_principle": 31014.898597926556
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 118,
            "towards_interest": 14162.06429645023,
            "towards_principle": 31199.695702072535
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 119,
            "towards_interest": 13976.166109558713,
            "towards_principle": 31385.593888964053
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 120,
            "towards_interest": 13789.160279303636,
            "towards_principle": 31572.599719219128
          }
        ],
        "year": 10
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 150471.1585360084,
        "annual_tax_saved": 45141.347560802526,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 121,
            "towards_interest": 13601.040205976624,
            "towards_principle": 31760.71979254614
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 122,
            "towards_interest": 13411.799250546033,
            "towards_principle": 31949.96074797673
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 123,
            "towards_interest": 13221.430734422673,
            "towards_principle": 32140.32926410009
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 124,
            "towards_interest": 13029.927939224075,
            "towards_principle": 32331.83205929869
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 125,
            "towards_interest": 12837.284106537421,
            "towards_principle": 32524.47589198534
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 126,
            "towards_interest": 12643.492437681009,
            "towards_principle": 32718.267560841756
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 127,
            "towards_interest": 12448.546093464327,
            "towards_principle": 32913.21390505844
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 128,
            "towards_interest": 12252.438193946688,
            "towards_principle": 33109.32180457607
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 129,
            "towards_interest": 12055.16181819442,
            "towards_principle": 33306.59818032834
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 130,
            "towards_interest": 11856.710004036631,
            "towards_principle": 33505.04999448614
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 131,
            "towards_interest": 11657.075747819485,
            "towards_principle": 33704.68425070328
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 132,
            "towards_interest": 11456.252004159045,
            "towards_principle": 33905.50799436372
          }
        ],
        "year": 11
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 121367.99628264687,
        "annual_tax_saved": 36410.398884794056,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 133,
            "towards_interest": 11254.231685692626,
            "towards_principle": 34107.52831283014
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 134,
            "towards_interest": 11051.00766282868,
            "towards_principle": 34310.75233569408
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 135,
            "towards_interest": 10846.57276349517,
            "towards_principle": 34515.187235027595
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 136,
            "towards_interest": 10640.919772886466,
            "towards_principle": 34720.8402256363
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 137,
            "towards_interest": 10434.041433208715,
            "towards_principle": 34927.71856531405
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 138,
            "towards_interest": 10225.930443423718,
            "towards_principle": 35135.829555099044
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 139,
            "towards_interest": 10016.579458991253,
            "towards_principle": 35345.18053953151
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 140,
            "towards_interest": 9805.981091609878,
            "towards_principle": 35555.778906912885
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 141,
            "towards_interest": 9594.12790895619,
            "towards_principle": 35767.63208956657
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 142,
            "towards_interest": 9381.01243442252,
            "towards_principle": 35980.747564100246
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 143,
            "towards_interest": 9166.62714685309,
            "towards_principle": 36195.132851669674
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 144,
            "towards_interest": 8950.96448027856,
            "towards_principle": 36410.7955182442
          }
        ],
        "year": 12
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 90114.3931831863,
        "annual_tax_saved": 27034.317954955888,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 145,
            "towards_interest": 8734.01682364902,
            "towards_principle": 36627.74317487374
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 146,
            "towards_interest": 8515.7765205654,
            "towards_principle": 36845.983477957365
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 147,
            "towards_interest": 8296.235869009237,
            "towards_principle": 37065.52412951353
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 148,
            "towards_interest": 8075.387121070886,
            "towards_principle": 37286.37287745188
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 149,
            "towards_interest": 7853.2224826760685,
            "towards_principle": 37508.53751584669
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 150,
            "towards_interest": 7629.734113310815,
            "towards_principle": 37732.02588521195
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 151,
            "towards_interest": 7404.914125744761,
            "towards_principle": 37956.845872778
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 152,
            "towards_interest": 7178.754585752791,
            "towards_principle": 38183.00541276997
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 153,
            "towards_interest": 6951.247511835037,
            "towards_principle": 38410.51248668773
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 154,
            "towards_interest": 6722.38487493519,
            "towards_principle": 38639.375123587575
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 155,
            "towards_interest": 6492.158598157146,
            "towards_principle": 38869.60140036562
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 156,
            "towards_interest": 6260.560556479969,
            "towards_principle": 39101.199442042795
          }
        ],
        "year": 13
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 56551.45255855411,
        "annual_tax_saved": 16965.435767566232,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 157,
            "towards_interest": 6027.58257647113,
            "towards_principle": 39334.17742205164
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 158,
            "towards_interest": 5793.2164359980725,
            "towards_principle": 39568.54356252469
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 159,
            "towards_interest": 5557.45386393803,
            "towards_principle": 39804.30613458474
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 160,
            "towards_interest": 5320.286539886129,
            "towards_principle": 40041.47345863664
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 161,
            "towards_interest": 5081.7060938617515,
            "towards_principle": 40280.053904661014
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 162,
            "towards_interest": 4841.704106013147,
            "towards_principle": 40520.05589250962
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 163,
            "towards_interest": 4600.272106320277,
            "towards_principle": 40761.48789220249
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 164,
            "towards_interest": 4357.4015742959045,
            "towards_principle": 41004.35842422686
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 165,
            "towards_interest": 4113.083938684886,
            "towards_principle": 41248.676059837875
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 166,
            "towards_interest": 3867.3105771616847,
            "towards_principle": 41494.44942136108
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 167,
            "towards_interest": 3620.072816026075,
            "towards_principle": 41741.68718249669
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 168,
            "towards_interest": 3371.3619298970325,
            "towards_principle": 41990.398068625735
          }
        ],
        "year": 14
      },
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 20508.536809360725,
        "annual_tax_saved": 6152.561042808217,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 169,
            "towards_interest": 3121.169141404804,
            "towards_principle": 42240.59085711796
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 170,
            "towards_interest": 2869.485620881143,
            "towards_principle": 42492.274377641625
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 171,
            "towards_interest": 2616.302486047695,
            "towards_principle": 42745.45751247507
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 172,
            "towards_interest": 2361.610801702531,
            "towards_principle": 43000.14919682023
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 173,
            "towards_interest": 2105.4015794048105,
            "towards_principle": 43256.358419117954
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 174,
            "towards_interest": 1847.665777157566,
            "towards_principle": 43514.0942213652
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 175,
            "towards_interest": 1588.3942990885987,
            "towards_principle": 43773.365699434165
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 176,
            "towards_interest": 1327.57799512947,
            "towards_principle": 44034.182003393296
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 177,
            "towards_interest": 1065.207660692585,
            "towards_principle": 44296.55233783018
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 178,
            "towards_interest": 801.2740363463469,
            "towards_principle": 44560.48596217642
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 179,
            "towards_interest": 535.7678074883789,
            "towards_principle": 44825.992191034384
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 180,
            "towards_interest": 268.6796040167991,
            "towards_principle": 45093.08039450597
          }
        ],
        "year": 15
      }
    ],
    "annual_emi_percentage_hike": 0,
    "total_interest": 3165110,
    "total_tax_saved": 949529,
    "total_years_to_settle": 15
  },
  {
    "annual_details": [
      {
        "annual_emi_amount": 544341.1199822732,
        "annual_interest": 351253.80645822175,
        "annual_tax_saved": 105376.14193746653,
        "monthly_emis": [
          {
            "emi_amount": 45361.759998522764,
            "month": 1,
            "towards_interest": 29791.666666666668,
            "towards_principle": 15570.093331856096
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 2,
            "towards_interest": 29698.89486056436,
            "towards_principle": 15662.865137958404
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 3,
            "towards_interest": 29605.570289117357,
            "towards_principle": 15756.189709405407
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 4,
            "towards_interest": 29511.68965876549,
            "towards_principle": 15850.070339757276
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 5,
            "towards_interest": 29417.249656324435,
            "towards_principle": 15944.51034219833
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 6,
            "towards_interest": 29322.246948868837,
            "towards_principle": 16039.513049653928
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 7,
            "towards_interest": 29226.678183614647,
            "towards_principle": 16135.081814908117
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 8,
            "towards_interest": 29130.53998780082,
            "towards_principle": 16231.220010721943
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 9,
            "towards_interest": 29033.82896857027,
            "towards_principle": 16327.931029952495
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 10,
            "towards_interest": 28936.541712850132,
            "towards_principle": 16425.218285672632
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 11,
            "towards_interest": 28838.674787231335,
            "towards_principle": 16523.08521129143
          },
          {
            "emi_amount": 45361.759998522764,
            "month": 12,
            "towards_interest": 28740.2247378474,
            "towards_principle": 16621.535260675366
          }
        ],
        "year": 1
      },
      {
        "annual_emi_amount": 571558.1759813867,
        "annual_interest": 336076.65053709317,
        "annual_tax_saved": 100822.99516112795,
        "monthly_emis": [
          {
            "emi_amount": 47629.8479984489,
            "month": 13,
            "towards_interest": 28641.18809025254,
            "towards_principle": 18988.65990819636
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 14,
            "towards_interest": 28528.047324966206,
            "towards_principle": 19101.800673482696
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 15,
            "towards_interest": 28414.2324292867,
            "towards_principle": 19215.6155691622
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 16,
            "towards_interest": 28299.739386520443,
            "towards_principle": 19330.10861192846
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 17,
            "towards_interest": 28184.56415604104,
            "towards_principle": 19445.28384240786
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 18,
            "towards_interest": 28068.702673146687,
            "towards_principle": 19561.145325302215
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 19,
            "towards_interest": 27952.150848916765,
            "towards_principle": 19677.697149532138
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 20,
            "towards_interest": 27834.90457006747,
            "towards_principle": 19794.943428381433
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 21,
            "towards_interest": 27716.959698806695,
            "towards_principle": 19912.888299642207
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 22,
            "towards_interest": 27598.312072687997,
            "towards_principle": 20031.535925760905
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 23,
            "towards_interest": 27478.95750446367,
            "towards_principle": 20150.890493985233
          },
          {
            "emi_amount": 47629.8479984489,
            "month": 24,
            "towards_interest": 27358.891781937007,
            "towards_principle": 20270.956216511895
          }
        ],
        "year": 2
      },
      {
        "annual_emi_amount": 600136.0847804561,
        "annual_interest": 317721.48037067894,
        "annual_tax_saved": 95316.44411120367,
        "monthly_emis": [
          {
            "emi_amount": 50011.34039837135,
            "month": 25,
            "towards_interest": 27238.110667813624,
            "towards_principle": 22773.229730557727
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 26,
            "towards_interest": 27102.420174002382,
            "towards_principle": 22908.92022436897
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 27,
            "towards_interest": 26965.921190998855,
            "towards_principle": 23045.419207372495
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 28,
            "towards_interest": 26828.608901554926,
            "towards_principle": 23182.731496816425
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 29,
            "towards_interest": 26690.47845971973,
            "towards_principle": 23320.861938651622
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 30,
            "towards_interest": 26551.524990668593,
            "towards_principle": 23459.815407702758
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 31,
            "towards_interest": 26411.743590531034,
            "towards_principle": 23599.596807840317
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 32,
            "towards_interest": 26271.12932621765,
            "towards_principle": 23740.2110721537
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 33,
            "towards_interest": 26129.677235246072,
            "towards_principle": 23881.66316312528
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 34,
            "towards_interest": 25987.38232556578,
            "towards_principle": 24023.95807280557
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 35,
            "towards_interest": 25844.23957538198,
            "towards_principle": 24167.10082298937
          },
          {
            "emi_amount": 50011.34039837135,
            "month": 36,
            "towards_interest": 25700.24393297833,
            "towards_principle": 24311.09646539302
          }
        ],
        "year": 3
      },
      {
        "annual_emi_amount": 630142.889019479,
        "annual_interest": 295850.64309613267,
        "annual_tax_saved": 88755.1929288398,
        "monthly_emis": [
          {
            "emi_amount": 52511.90741828992,
            "month": 37,
            "towards_interest": 25555.3903165387,
            "towards_principle": 26956.517101751222
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 38,
            "towards_interest": 25394.77440214076,
            "towards_principle": 27117.13301614916
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 39,
            "towards_interest": 25233.201484586207,
            "towards_principle": 27278.705933703714
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 40,
            "towards_interest": 25070.665861731217,
            "towards_principle": 27441.241556558703
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 41,
            "towards_interest": 24907.161797456723,
            "towards_principle": 27604.745620833197
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 42,
            "towards_interest": 24742.683521465922,
            "towards_principle": 27769.223896823998
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 43,
            "towards_interest": 24577.225229080683,
            "towards_principle": 27934.682189209238
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 44,
            "towards_interest": 24410.781081036643,
            "towards_principle": 28101.126337253278
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 45,
            "towards_interest": 24243.345203277175,
            "towards_principle": 28268.562215012746
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 46,
            "towards_interest": 24074.911686746054,
            "towards_principle": 28436.995731543866
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 47,
            "towards_interest": 23905.474587178942,
            "towards_principle": 28606.43283111098
          },
          {
            "emi_amount": 52511.90741828992,
            "month": 48,
            "towards_interest": 23735.02792489357,
            "towards_principle": 28776.87949339635
          }
        ],
        "year": 4
      },
      {
        "annual_emi_amount": 661650.033470453,
        "annual_interest": 270096.3951607733,
        "annual_tax_saved": 81028.91854823199,
        "monthly_emis": [
          {
            "emi_amount": 55137.502789204416,
            "month": 49,
            "towards_interest": 23563.565684578753,
            "towards_principle": 31573.937104625664
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 50,
            "towards_interest": 23375.437642663688,
            "towards_principle": 31762.065146540728
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 51,
            "towards_interest": 23186.188671165553,
            "towards_principle": 31951.314118038863
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 52,
            "towards_interest": 22995.81209121224,
            "towards_principle": 32141.690697992177
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 53,
            "towards_interest": 22804.3011841367,
            "towards_principle": 32333.201605067716
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 54,
            "towards_interest": 22611.649191239838,
            "towards_principle": 32525.85359796458
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 55,
            "towards_interest": 22417.849313551967,
            "towards_principle": 32719.65347565245
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 56,
            "towards_interest": 22222.89471159287,
            "towards_principle": 32914.60807761154
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 57,
            "towards_interest": 22026.778505130434,
            "towards_principle": 33110.72428407398
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 58,
            "towards_interest": 21829.493772937825,
            "towards_principle": 33308.00901626659
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 59,
            "towards_interest": 21631.03355254924,
            "towards_principle": 33506.469236655175
          },
          {
            "emi_amount": 55137.502789204416,
            "month": 60,
            "towards_interest": 21431.39084001417,
            "towards_principle": 33706.11194919025
          }
        ],
        "year": 5
      },
      {
        "annual_emi_amount": 694732.5351439756,
        "annual_interest": 240058.42184130766,
        "annual_tax_saved": 72017.5265523923,
        "monthly_emis": [
          {
            "emi_amount": 57894.37792866464,
            "month": 61,
            "towards_interest": 21230.558589650245,
            "towards_principle": 36663.81933901439
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 62,
            "towards_interest": 21012.103332755283,
            "towards_principle": 36882.27459590936
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 63,
            "towards_interest": 20792.346446621326,
            "towards_principle": 37102.03148204331
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 64,
            "towards_interest": 20571.28017570748,
            "towards_principle": 37323.097752957154
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 65,
            "towards_interest": 20348.89671826278,
            "towards_principle": 37545.481210401864
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 66,
            "towards_interest": 20125.188226050803,
            "towards_principle": 37769.18970261383
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 67,
            "towards_interest": 19900.14680407273,
            "towards_principle": 37994.23112459191
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 68,
            "towards_interest": 19673.7645102887,
            "towards_principle": 38220.61341837594
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 69,
            "towards_interest": 19446.033355337542,
            "towards_principle": 38448.34457332709
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 70,
            "towards_interest": 19216.945302254808,
            "towards_principle": 38677.432626409834
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 71,
            "towards_interest": 18986.49226618911,
            "towards_principle": 38907.88566247553
          },
          {
            "emi_amount": 57894.37792866464,
            "month": 72,
            "towards_interest": 18754.66611411686,
            "towards_principle": 39139.71181454778
          }
        ],
        "year": 6
      },
      {
        "annual_emi_amount": 729469.1619011746,
        "annual_interest": 205301.16062606825,
        "annual_tax_saved": 61590.34818782047,
        "monthly_emis": [
          {
            "emi_amount": 60789.09682509788,
            "month": 73,
            "towards_interest": 18521.45866455518,
            "towards_principle": 42267.6381605427
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 74,
            "towards_interest": 18269.613987181947,
            "towards_principle": 42519.48283791593
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 75,
            "towards_interest": 18016.268735272697,
            "towards_principle": 42772.82808982518
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 76,
            "towards_interest": 17761.41396790416,
            "towards_principle": 43027.68285719372
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 77,
            "towards_interest": 17505.040690880047,
            "towards_principle": 43284.056134217826
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 78,
            "towards_interest": 17247.13985641366,
            "towards_principle": 43541.956968684215
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 79,
            "towards_interest": 16987.702362808584,
            "towards_principle": 43801.39446228929
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 80,
            "towards_interest": 16726.719054137448,
            "towards_principle": 44062.37777096043
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 81,
            "towards_interest": 16464.180719918808,
            "towards_principle": 44324.91610517907
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 82,
            "towards_interest": 16200.078094792114,
            "towards_principle": 44589.01873030576
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 83,
            "towards_interest": 15934.401858190711,
            "towards_principle": 44854.69496690716
          },
          {
            "emi_amount": 60789.09682509788,
            "month": 84,
            "towards_interest": 15667.14263401289,
            "towards_principle": 45121.95419108499
          }
        ],
        "year": 7
      },
      {
        "annual_emi_amount": 765942.6199962333,
        "annual_interest": 165350.9133249719,
        "annual_tax_saved": 49605.27399749157,
        "monthly_emis": [
          {
            "emi_amount": 63828.551666352774,
            "month": 85,
            "towards_interest": 15398.290990291009,
            "towards_principle": 48430.260676061764
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 86,
            "towards_interest": 15109.727353762808,
            "towards_principle": 48718.82431258997
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 87,
            "towards_interest": 14819.444358900291,
            "towards_principle": 49009.10730745248
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 88,
            "towards_interest": 14527.431761193384,
            "towards_principle": 49301.11990515939
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 89,
            "towards_interest": 14233.679255091813,
            "towards_principle": 49594.87241126096
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 90,
            "towards_interest": 13938.176473641382,
            "towards_principle": 49890.37519271139
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 91,
            "towards_interest": 13640.912988118142,
            "towards_principle": 50187.63867823463
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 92,
            "towards_interest": 13341.878307660327,
            "towards_principle": 50486.673358692446
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 93,
            "towards_interest": 13041.06187889812,
            "towards_principle": 50787.48978745466
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 94,
            "towards_interest": 12738.4530855812,
            "towards_principle": 51090.09858077158
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 95,
            "towards_interest": 12434.041248204105,
            "towards_principle": 51394.51041814867
          },
          {
            "emi_amount": 63828.551666352774,
            "month": 96,
            "towards_interest": 12127.815623629302,
            "towards_principle": 51700.736042723474
          }
        ],
        "year": 8
      },
      {
        "annual_emi_amount": 804239.750996045,
        "annual_interest": 119692.73062141852,
        "annual_tax_saved": 35907.81918642556,
        "monthly_emis": [
          {
            "emi_amount": 67019.97924967042,
            "month": 97,
            "towards_interest": 11819.765404708074,
            "towards_principle": 55200.213844962345
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 98,
            "towards_interest": 11490.864130548505,
            "towards_principle": 55529.11511912191
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 99,
            "towards_interest": 11160.003152963738,
            "towards_principle": 55859.97609670668
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 100,
            "towards_interest": 10827.170795387527,
            "towards_principle": 56192.80845428289
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 101,
            "towards_interest": 10492.35531168076,
            "towards_principle": 56527.62393798966
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 102,
            "towards_interest": 10155.544885716903,
            "towards_principle": 56864.434363953515
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 103,
            "towards_interest": 9816.727630965015,
            "towards_principle": 57203.2516187054
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 104,
            "towards_interest": 9475.891590070229,
            "towards_principle": 57544.08765960019
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 105,
            "towards_interest": 9133.024734431776,
            "towards_principle": 57886.95451523864
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 106,
            "towards_interest": 8788.114963778478,
            "towards_principle": 58231.86428589194
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 107,
            "towards_interest": 8441.150105741706,
            "towards_principle": 58578.82914392871
          },
          {
            "emi_amount": 67019.97924967042,
            "month": 108,
            "towards_interest": 8092.117915425797,
            "towards_principle": 58927.86133424462
          }
        ],
        "year": 9
      },
      {
        "annual_emi_amount": 844451.7385458476,
        "annual_interest": 67767.05154324861,
        "annual_tax_saved": 20330.115462974583,
        "monthly_emis": [
          {
            "emi_amount": 70370.97821215395,
            "month": 109,
            "towards_interest": 7741.006074975921,
            "towards_principle": 62629.972137178025
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 110,
            "towards_interest": 7367.835824325235,
            "towards_principle": 63003.14238782871
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 111,
            "towards_interest": 6992.442100931089,
            "towards_principle": 63378.53611122286
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 112,
            "towards_interest": 6614.81165660172,
            "towards_principle": 63756.16655555223
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 113,
            "towards_interest": 6234.93116420822,
            "towards_principle": 64136.04704794573
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 114,
            "towards_interest": 5852.7872172142115,
            "towards_principle": 64518.19099493974
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 115,
            "towards_interest": 5468.366329202696,
            "towards_principle": 64902.61188295125
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 116,
            "towards_interest": 5081.654933400111,
            "towards_principle": 65289.323278753836
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 117,
            "towards_interest": 4692.639382197536,
            "towards_principle": 65678.33882995641
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 118,
            "towards_interest": 4301.3059466690465,
            "towards_principle": 66069.6722654849
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 119,
            "towards_interest": 3907.6408160871983,
            "towards_principle": 66463.33739606675
          },
          {
            "emi_amount": 70370.97821215395,
            "month": 120,
            "towards_interest": 3511.6300974356336,
            "towards_principle": 66859.34811471832
          }
        ],
        "year": 10
      }
    ],
    "annual_emi_percentage_hike": 5,
    "total_interest": 2369164,
    "total_tax_saved": 710746,
    "total_years_to_settle": 10
  }
]
```
</details>
<br />

## API Design and Documentation
Find the details of API Design and documentation on Swagger [here](https://app.swaggerhub.com/apis-docs/AxiomSamarth/home-loan-optimizer/1.0.0#/)

## Current assumptions in the algorithm
- Starting month of the loan is assumed to be April of the financial year.
- Only taxes saved on interest paid u/s 24 is calculated and compared.
- It is assumed that the borrower falls in the 30% tax slab of the old regime.

## Features to be added in future -

- Recommendation of alternative & parallel investment options to make the home loan effectively interest free
- Consideration of additional prepayments of loan during the recommendation calculation
- Consideration of sections 80C, 80EE, 80EEA during recommendation calculation

## Contribution guidelines - 

We are glad that you want to contribute to this project and make it better. Please refer to [contribution guidelines](./contributing.md) and follow the guidelines. 
