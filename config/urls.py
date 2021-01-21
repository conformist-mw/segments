from django.conf import settings
from django.contrib import admin
from django.urls import include, path

admin.site.site_header = 'Отрезки'

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('segments.urls')),
]

if settings.DEBUG:
    import debug_toolbar
    from django.conf.urls.static import static
    urlpatterns += [path('__debug__/', include(debug_toolbar.urls))]
    urlpatterns += static(
        settings.MEDIA_URL, document_root=settings.MEDIA_ROOT,
    )
