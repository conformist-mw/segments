from django import forms

from .models import Color, ColorType, Segment


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
