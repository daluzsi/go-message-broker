import http from 'k6/http';
import { check } from 'k6';
import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

const queues = [
  'my-queue-payment',
  'my-queue-invoice',
  'my-queue-deposit',
  'my-queue-withdrawal',
];

const baseURL = 'http://localhost.localstack.cloud:4566/000000000000/';

export const options = {
  stages: [
    { duration: '30s', target: 10 }, // 10 VUs durante 30 segundos
    { duration: '1m', target: 20 },  // 20 VUs durante 1 minuto
    { duration: '30s', target: 0 },  // Desacelera para 0 VUs
  ],
};

export default function () {
  queues.forEach(queue => {
    const url = `${baseURL}${queue}`;
    const params = {
      headers: {
        'Accept': 'application/json',
      },
    };

    const messageBody = JSON.stringify({
      transaction_id: Math.random().toString(36).substring(2, 15),
      amount: randomIntBetween(1, 1000), // valor aleatório entre 1 e 1000
      currency: 'USD',
      timestamp: new Date().toISOString()
    });

    const res = http.post(`${url}?Action=SendMessage&MessageBody=${encodeURIComponent(messageBody)}`, null, params);

    check(res, {
      'status was 200': (r) => r.status === 200,
    });
  });
}