from django import forms

from .models import Color, ColorType, Rack, Segment


class SegmentCreateForm(forms.ModelForm):

    color_type = forms.ModelChoiceField(
        queryset=ColorType.objects.all(),
        to_field_name='name',
    )
    color = forms.ModelChoiceField(
        queryset=Color.objects.all(),
        to_field_name='name',
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
        choices=[('0', 'Все')] + [(ct.id, ct.name) for ct in ColorType.objects.all()],
        required=False,
        widget=forms.Select(attrs={'class': 'form-control'})
    )
    color = forms.ChoiceField(
        choices=[('0', 'Все')] + [(c.id, c.name) for c in Color.objects.all()],
        required=False,
    )
    width = forms.IntegerField(required=False)
    height = forms.IntegerField(required=False)
    deleted = forms.BooleanField(required=False)
