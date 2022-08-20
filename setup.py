from setuptools import find_packages, setup
from bupt_session_py import __about__, __author__, __version__

# read the contents of your README file
from os import path
this_directory = path.abspath(path.dirname(__file__))
with open(path.join(this_directory, 'README-py.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
    name='bupt_session',
    version=__version__,
    author=__author__,
    author_email='rinchannow@bupt.edu.cn',
    description=__about__,
    long_description=long_description,
    long_description_content_type='text/markdown',
    classifiers=[
        'Development Status :: 5 - Production/Stable',
        'Intended Audience :: Developers',
        'Topic :: Software Development :: Libraries',
        'Programming Language :: Python :: 3',
    ],
    keywords='python BUPT session login',
    url='https://github.com/OpenBUPT/session',
    license='MIT',
    packages=find_packages(),
    include_package_data=True,
    zip_safe=True,
    install_requires=open('requirements.txt').read().split(),
)
