```text
                                      marshal             unmarshal           alloc (unmarshal)
timeseries - small        vtproto     2060 ns/op          10095 ns/op         27 kB/op
                          csproto     4159 ns/op          12687 ns/op         22 kB/op
                          csproto*    4161 ns/op          10186 ns/op         18 kB/op
                          proto       5831 ns/op          14592 ns/op         22 kB/op
timeseries - medium       vtproto     195921 ns/op        1070954 ns/op       2.6 MB/op
                          csproto     408949 ns/op        1364211 ns/op       2.5 MB/op
                          csproto*    411395 ns/op        1270685 ns/op       2.2 MB/op
                          proto       594622 ns/op        1597091 ns/op       2.5 MB/op
timeseries - large        vtproto     20523750 ns/op      98311750 ns/op      271 MB/op
                          csproto     40058865 ns/op      117696310 ns/op     259 MB/op
                          csproto*    40513806 ns/op      93412525 ns/op      227 MB/op
                          proto       56011681 ns/op      138445359 ns/op     259 MB/op
address book - small      vtproto     5578 ns/op          24878 ns/op         40 kB/op
                          csproto     7145 ns/op          28965 ns/op         40 kB/op
                          csproto*    7046 ns/op          22451 ns/op         30 kB/op
                          proto       17670 ns/op         40185 ns/op         41 kB/op
address book - medium     vtproto     914181 ns/op        5680822 ns/op       9.5 MB/op
                          csproto     1069070 ns/op       6744685 ns/op       9.5 MB/op
                          csproto*    1025744 ns/op       5002779 ns/op       7.3 MB/op
                          proto       2457689 ns/op       8099329 ns/op       9.5 MB/op
address book - large      vtproto     121558263 ns/op     426731444 ns/op     954 MB/op
                          csproto     114333762 ns/op     510945166 ns/op     954 MB/op 
                          csproto*    110876516 ns/op     356000472 ns/op     729 MB/op
                          proto       241064441 ns/op     659395375 ns/op     954 MB/op
```
