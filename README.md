# Benchmarking

## Concurrent 100 client; total 10,000 request
<img width="1258" alt="100-user-100-cycle" src="https://user-images.githubusercontent.com/19304394/100639627-a0588e80-335f-11eb-91bc-72b1ff373730.png">

## Concurrent 200 client; total 20,000 request
<img width="1261" alt="200-user-100-cycle" src="https://user-images.githubusercontent.com/19304394/100639619-9d5d9e00-335f-11eb-8e35-141e30047a97.png">

**P.S** 

_The error percentage is not 0 in the second image. It's because of the `ulimit`. This benchmarking was done in Mac with `ulimit` 4096. This is not enough to handle all those concurrent request. But as you can see there was no error(TCP ) in the first image with 100 concurrent clients doing 10K requests. So it meanse if we run this test in a Linux server with `ulimit` 65,536 (this is the default one) we more likely won't observe any _