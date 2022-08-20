from typing_extensions import Self
import requests
import re
from bupt_session_py.api import LOGIN

from bupt_session_py.exceptions import LoginFailed


class Session(requests.Session):

    def __init__(self) -> None:
        super().__init__()
        # The default header.
        self.headers.update(
            {"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36"})

    def login(self, username: str, password: str) -> Self:
        """登录到统一网关, 将登录信息维护在 Session 中。

        Parameters
        ----------
        username : str 北邮学工号
        password : str 信息门户密码

        Returns
        -------
        Session: 当前 Session

        Raises
        ------
        LoginFailed: 登录失败
        """

        res = self.get(LOGIN)
        execution = re.findall(
            r'input name="execution" value="(.*)"/><input name="_eventId"', str(res.content))
        if len(execution) == 0:
            raise LoginFailed('Execution code is not found.')
        execution = execution[0]

        login_form = {
            'submit': "LOGIN",
            'type': "username_password",
            '_eventId': "submit",
            'username': username,
            'password': password,
            'execution': execution,
        }

        # get login cookies
        res = self.post(
            LOGIN, data=login_form, allow_redirects=False)
        if res.status_code != requests.codes.found:  # 302 is the expected code
            raise LoginFailed(f'302 is expected, but get {res.status_code}')

        return self
