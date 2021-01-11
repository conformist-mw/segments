from django.http import JsonResponse
from django.views.generic import CreateView, ListView

from .forms import SegmentCreateForm
from .models import ColorType, Rack, Segment


class SegmentsListView(ListView):
    model = Segment
    template_name = 'segments.html'
    context_object_name = 'segments'
    queryset = Segment.objects.select_related('color', 'rack', 'color__type')

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        colors = {
            ct.name: [c.name for c in ct.colors.all()]
            for ct in ColorType.objects.prefetch_related('colors').all()
        }
        context.update({
            'form': SegmentCreateForm(),
            'colors': colors,
            'racks': {r.name: r.id for r in Rack.objects.order_by('name')},
        })
        return context


class SegmentCreateView(CreateView):
    form_class = SegmentCreateForm

    def form_valid(self, form):
        form.save()
        return JsonResponse({}, status=201)

    def form_invalid(self, form):
        return JsonResponse(form.errors.as_json(), status=400, safe=False)
