import os

from django.conf import settings
from django.contrib import admin
from django.urls import include, path

from segments.forms import CustomAdminAuthenticationForm

version = os.getenv('HEROKU_RELEASE_VERSION', 'dev')

admin.site.site_header = f'Отрезки {version}'
admin.site.login_form = CustomAdminAuthenticationForm

urlpatterns = [
    path('admin/', admin.site.urls),
    path('api/', include('api.urls')),
    path('', include('segments.urls')),
]

if settings.DEBUG:
    import debug_toolbar
    from django.conf.urls.static import static

    # have to insert it because it was intercepted by nested <slug><slug> url
    urlpatterns = [
        path('__debug__/', include(debug_toolbar.urls)),
    ] + urlpatterns
    urlpatterns += static(
        settings.MEDIA_URL, document_root=settings.MEDIA_ROOT,
    )
