# xk6-fifo

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/grafana/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
|------|

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```shell
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  $ xk6 build --with github.com/nj-apps/xk6-fifo@latest
  ```

## Example

```javascript
// script.js
import fifo from 'k6/x/fifo';

export const options = {
  scenarios: {
    generator: {
      exec: 'generator',
      executor: 'per-vu-iterations',
      vus: 2,
    },
    results: {
      exec: 'consumer',
      executor: 'per-vu-iterations',
      startTime: '1s',
      maxDuration: '2s',
      vus: 1,
        iterations: 7
    },
  },
};

const client = new fifo.Client();

export function generator() {
  client.push(`hello_${__VU}_1`);
  client.push(`hello_${__VU}_2`);
  client.push(`hello_${__VU}_3`);
}

export function consumer() {
    try{
        console.log(client.pop());
    }
    catch(err){
        console.log("No more data");
    }
}

```

Result output:

```
$ ./k6 run script.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: script.js
     output: -

  scenarios: (100.00%) 2 scenarios, 3 max VUs, 10m30s max duration (incl. graceful stop):
           * generator: 1 iterations for each of 2 VUs (maxDuration: 10m0s, exec: generator, gracefulStop: 30s)
           * results: 7 iterations for each of 1 VUs (maxDuration: 2s, exec: consumer, startTime: 1s, gracefulStop: 30s)

INFO[0001] hello_1_1                                     source=console
INFO[0001] hello_1_2                                     source=console
INFO[0001] hello_1_3                                     source=console
INFO[0001] hello_3_1                                     source=console
INFO[0001] hello_3_2                                     source=console
INFO[0001] hello_3_3                                     source=console
INFO[0001] No more data                                  source=console

     data_received........: 0 B 0 B/s
     data_sent............: 0 B 0 B/s
     iteration_duration...: avg=135.14µs min=34.64µs med=105.74µs max=311.78µs p(90)=267.06µs p(95)=289.42µs
     iterations...........: 9   8.979212/s
     vus..................: 0   min=0      max=0
     vus_max..............: 3   min=3      max=3


running (00m01.0s), 0/3 VUs, 9 complete and 0 interrupted iterations
generator ✓ [======================================] 2 VUs  00m00.0s/10m0s  2/2 iters, 1 per VU
results   ✓ [======================================] 1 VUs  0.0s/2s         7/7 iters, 7 per VU

```

You can name the FIFOs if you need to use more than one in a script :

```javascript
const client1 = new fifo.Client("FIRST_FIFO");
const client2 = new fifo.Client("ANOTHER_FIFO");
```