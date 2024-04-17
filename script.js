import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  // A number specifying the number of VUs to run concurrently.
  vus: 400,
  // A string specifying the total duration of the test run.
  duration: '60s',

}
export default function() {
  const payload = JSON.stringify({
    latitude: 0.0,
    longitude: 0.0,
    nearest: 5
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };
  http.post("http://localhost:8080/ngo/nearest", payload, params)
  sleep(1);
}
