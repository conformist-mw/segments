from django.http import JsonResponse
from django.shortcuts import render
from django.views.generic import CreateView, ListView, UpdateView, View

from .forms import PrintSegmentsForm, SegmentCreateForm
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
            'print_form': PrintSegmentsForm(),
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
        return JsonResponse(
            {'error': form.errors.as_json()}, status=400, safe=False,
        )


class PrintSegmentsView(View):
    template_name = 'table.html'
    form_class = PrintSegmentsForm

    def post(self, request, *args, **kwargs):
        form = self.form_class(request.POST)
        if not form.is_valid():
            return JsonResponse(form.errors.as_json(), status=400, safe=False)
        print_rack = form.cleaned_data['print_rack']
        qs = Segment.objects.select_related('color', 'color__type', 'rack')
        if print_rack:
            qs = qs.filter(rack=print_rack)
        return render(request, self.template_name, {'segments': qs})


class MoveSegmentView(UpdateView):
    model = Segment
    fields = ['rack']

    def form_valid(self, form):
        segment = form.save()
        return JsonResponse({'rack': segment.rack.name}, status=200)

    def form_invalid(self, form):
        return JsonResponse({}, status=400)
