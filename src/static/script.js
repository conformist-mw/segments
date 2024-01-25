'use strict';

function filterColors() {
  var type = $(this).val();
  var colorSelect = $(this).parents('.colors').find('select').last();
  colorSelect.find('option:selected').prop('selected', false);
  if (!!type) {
    colorSelect.find(`option[data-type="${type}"]`).show();
    colorSelect.find(`option:not(option[data-type="${type}"])`).hide();
  } else {
    colorSelect.find('option').show();
  }
  colorSelect.focus();
};

function setColorType() {
  var type = $(this).find('option:selected').data('type');
  $(this).parents('.colors').find('select').first().find(`option[value="${type}"]`).prop('selected', true);
};

// add segments form
$('#add').submit(function(e) {
    e.preventDefault();
    const form = $(this);
    form.find('.text-danger').remove();
    $.ajax({
        type: 'POST',
        url: `${window.location.pathname}/add`,
        data: form.serialize(),
        success: function() {
            location.reload();
        },
        error: function(data) {
            const errors = JSON.parse(data.responseJSON.error);
            $.each(errors, function(field, error) {
                const small = $('<small />').attr('class', 'text-danger').html(error[0].message);
                $(`#id_${field}`).parent().append(small);
            });
        }
    });
});

// move segments to another rack form
$('button.edit').click(function() {
    $(this).parents('div').next().show();
});
$('button.move').click(function(e) {
    e.preventDefault();
    const segment_id = $(this).val();
    const form = $(this).parent();
    const rack = $(this).parent().prev().find('.rack').show();
    const newName = form.find("input[type='radio']:checked").parent().text();
    $.ajax({
        url: `${window.location.pathname}/move/${segment_id}`,
        data: form.serialize(),
        type: 'POST',
        success: function(result){
            rack.find('strong').text(newName);
        },
        complete: function() {
            form.hide();
        }
    });
});

// print segments
$('a.print').on('click', function() {
    $('div.print-form').toggle('slow');
});
$('#print-form').on('submit', function(e) {
    e.preventDefault();
    const csrftoken = Cookies.get('csrftoken');
    const form = $(this);
    $.ajax({
        headers: { 'X-CSRFToken': csrftoken },
        type: 'POST',
        url: `${window.location.pathname}print/`,
        data: form.serialize(),
        success: function(result){
            const new_window = window.open('', 'new_window', 'status=1');
            new_window.document.write(result);
        },
    });
});

// remove segment
$('.remove-toggle').click(function(){
  const segmentId = $(this).val();
  const removeElem = $('div[data-toggle="' + segmentId + '"]');
  $('div.remove-form').not(removeElem).each(function(){
    $(this).hide('fast');
  });
  removeElem.toggle('slow');
});
$('.removeSegment').click(function(e){
  e.preventDefault();
  const csrftoken = Cookies.get('csrftoken');
  const parent = $(this).parents('div.parent');
  const segmentId = $(this).val();
  const form = $(this).parents('form');
  $.ajax({
    headers: { 'X-CSRFToken': csrftoken },
    url: `${window.location.pathname}remove/${segmentId}`,
    data: form.serialize(),
    type: 'post',
    success: function(result){
      parent.remove();
    },
    error: function(error){
      form.find('small').text(error.responseJSON.message);
    }
  });
});

// activate segment form
$('.activate').click(function(){
  const segmentId = $(this).val();
  const parent = $(this).parents('div.parent');
  $.ajax({
    url: `${window.location.pathname}/activate/${segmentId}`,
    type: 'post',
    success: function(result){
      parent.remove();
    }
  })
})

// search segments form
$('#reset-form').on('click', function() {
  window.location.replace(window.location.pathname);
});
