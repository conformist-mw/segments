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
    urlpatterns += [path('__debug__/', include(debug_toolbar.urls))]
