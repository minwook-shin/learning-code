from django.contrib import admin
from .models import DB
from django.contrib.auth.models import Group
from django.utils.safestring import mark_safe

# Register your models here.

admin.site.site_header = "Custom Admin"
admin.site.site_title = "Custom Admin Portal"
admin.site.index_title = "Welcome to Custom Admin Portal"

admin.site.unregister(Group)


class MemberAdmin(admin.ModelAdmin):
    list_per_page = 5
    list_display = ('id', 'test', 'permission','content_size')
    list_editable = ('permission',)
    list_filter = ('permission',)
    search_fields = ('test',)
    ordering = ('id', 'test', 'permission',)

    fieldsets = (('기본', {'fields': (('test',),), },), ('범위', {'fields': (('permission',),), },))

    def mark_real(self, request, queryset):
        queryset.update(permission='real')

    mark_real.short_description = "test/real 범위 변경"
    actions = ["mark_real"]

    def content_size(self, text):
        return mark_safe('<u>{}</u>글자'.format(len(text.test)))

    content_size.short_description = '글자수'


admin.site.register(DB, MemberAdmin)
