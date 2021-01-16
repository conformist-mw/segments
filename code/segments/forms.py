from django import forms

from .models import Color, ColorType, Rack, Segment
from django.forms import renderers


class ColorSelect(forms.Select):
    option_template_name = 'forms/widgets/select_option.html'

    def create_option(
            self,
            name,
            value,
            label,
            selected,
            index,
            subindex=None,
            attrs=None,
    ):
        option = super().create_option(
            name, value, label, selected, index, subindex, attrs,
        )
        if value:
            color_type = label.split()[0]
            option.update({
                'color_type': color_type,
                'label': value,
            })
        return option


class SegmentCreateForm(forms.ModelForm):

    color_type = forms.ModelChoiceField(
        queryset=ColorType.objects.all(),
        to_field_name='name',
        empty_label='Фактура',
    )
    color = forms.ModelChoiceField(
        queryset=Color.objects.all(),
        to_field_name='name',
        empty_label='Выберите цвет',
        widget=ColorSelect(),
    )
    rack = forms.ModelChoiceField(
        queryset=Rack.objects.all(),
        empty_label='Расположение',
    )

    class Meta:
        model = Segment
        fields = ['color', 'width', 'height', 'rack', 'color_type']


class PrintSegmentsForm(forms.Form):
    print_rack = forms.ModelChoiceField(
        label='Расположение',
        queryset=Rack.objects.all(),
        empty_label='Все',
        required=False,
    )


class SearchSegmentsForm(forms.Form):

    color_type = forms.ChoiceField(
        choices=[('', 'Все')] + [(ct.name, ct.name) for ct in ColorType.objects.all()],
        required=False,
        widget=forms.Select(),
    )
    color = forms.ChoiceField(
        choices=[('', 'Все')] + [(c.name, str(c)) for c in Color.objects.all()],
        required=False,
        widget=ColorSelect(),
    )
    width = forms.IntegerField(required=False)
    height = forms.IntegerField(required=False)
    deleted = forms.BooleanField(required=False, label='Удалённые')
    order_number = forms.CharField(required=False)