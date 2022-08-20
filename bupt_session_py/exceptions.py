class LoginFailed(Exception):
    """登录失败"""

    def __init__(self, msg: str) -> None:
        if msg == '':
            super().__init__(f'Login failed')
        else:
            super().__init__(f'Login failed: {msg}')
