# BUPT Session

北邮统一登录网关 Session。用于需要登录的网络请求。

## Usage

```python
from bupt_session_py.session import Session

session = Session()

try:
    session.login('2022114514', 'yjsp1919810')

    # 进行一些北邮服务的网络请求
    # 这里以查询电费服务为例子 
    # (More detail: https://github.com/jerrymakesjelly/electricity-monitor)
    session.post('https://app.bupt.edu.cn/buptdf/wap/default/chong', allow_redirects=False)
    session.post('other url')
    # ...
except Exception as e:
    print(e)
```
