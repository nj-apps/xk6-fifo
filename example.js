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
