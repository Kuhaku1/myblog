from django.db import models
from markdownx.models import MarkdownxField
# Create your models here.


class userManager(models.Manager):
    def get_queryset(self):
        return super(userManager, self).get_queryset().filter(is_Delete=False)


class User(models.Model):
    """用户"""
    # id
    id = models.AutoField(primary_key=True)
    username = models.CharField(max_length=40, null=True,
                                blank=True, unique=True, verbose_name='用户名')
    password = models.CharField(max_length=40, null=True,
                                blank=True, verbose_name='密码')
    name = models.CharField(max_length=40, null=True,
                                blank=True, unique=True, verbose_name='名称')
    is_Delete = models.BooleanField(default=False, verbose_name='是否删除')
    # 管理器
    object = models.Manager()
    userManager = userManager()

    def __str__(self):
        return self.name

    class Meta:
        db_table = 'User'
        verbose_name = '用户'
        verbose_name_plural = '用户'

class Article(models.Model):
    """文章"""
    # id
    id = models.AutoField(primary_key=True)
    myfield = MarkdownxField(verbose_name='内容')
    
    # 管理器
    object = models.Manager()

    def __str__(self):
        return "lixingyu"

    class Meta:
        db_table = 'Article'
        verbose_name = '文章'
        verbose_name_plural = '文章'
