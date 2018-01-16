
try:
    import urllib.parse as urlparse
except ImportError:
    import urlparse


class ArchitectException(Exception):
    pass


class ArchitectClient(object):

    def __init__(self, api_url='http://localhost:8181', inventory='default'):
        self.api_url = api_url
        self.inventory = inventory

    def _req_get(self, path):
        '''
        A thin wrapper to get http method of architect api
        print(api._req_get('/keys'))
        '''
        import requests

        headers = {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        }
        # self._ssl_verify = self.ignore_ssl_errors
        params = {'url': self._construct_url(path),
                  'headers': headers}
        try:
            resp = requests.get(**params)

            if resp.status_code == 401:
                raise ArchitectException(str(resp.status_code) + ':Authentication denied')
                return

            if resp.status_code == 500:
                raise ArchitectException('{}: Server error.'.format(resp.status_code))
                return

            if resp.status_code == 404:
                raise ArchitectException(str(resp.status_code) +' :This request returns nothing.')
                return
        except ArchitectException as e:
            print(e)
            return
        return resp.json()

    def _construct_url(self, path):
        '''
        Construct the url to architect-api for the given path
        Args:
            path: the path to the architect-api resource
        '''

        relative_path = path.lstrip('/')
        return urlparse.urljoin(self.api_url, relative_path)

    def get_data(self, source, resource=None):
        if resource is None:
            path = '/inventory/v1/{}/data.json?source={}'.format(self.inventory,
                                                                 source)
        else:
            path = '/inventory/v1/{}/{}/data.json?source={}'.format(self.inventory,
                                                                    resource,
                                                                    source)
        return self._req_get(path)
