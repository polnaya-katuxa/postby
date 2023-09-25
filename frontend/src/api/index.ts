import {
  Configuration, DefaultApi,
} from '@/openapi';
import Cookies from 'cookies-ts';

export default new DefaultApi(new Configuration({
  basePath: 'http://localhost:8088',
  baseOptions: {
    headers: {
      'User-Token': new Cookies().get('user-token'),
    },
  },
}));
