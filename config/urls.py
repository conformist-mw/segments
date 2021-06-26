from django.conf import settings
from django.contrib import admin
from django.urls import include, path

from segments.forms import CustomAdminAuthenticationForm

admin.site.site_header = 'Отрезки'
admin.site.login_form = CustomAdminAuthenticationForm

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('segments.urls')),
]

if settings.DEBUG:
    import debug_toolbar
    from django.conf.urls.static import static

    # have to insert it because it was intercepted by nested <slug><slug> url
    urlpatterns.insert(0, path('__debug__/', include(debug_toolbar.urls)))
    urlpatterns += static(
        settings.MEDIA_URL, document_root=settings.MEDIA_ROOT,
    )
