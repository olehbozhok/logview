
const filterInp = ` <span class=" input-group mb-12 mr-sm-12">
<div class="input-group-prepend" onclick="removeParent(this)">
    <div class="input-group-text">-</div>
</div>
<input type="text" class="form-control __tmpl_class__" placeholder="filter">
</span>`
const __tmpl_class__ = "__tmpl_class__"

const filterClass = "filterClass"

function addFilterInput(selector, identityClass) {
    $(selector).append(filterInp.replace(__tmpl_class__, identityClass))
}

function removeParent(el) {
    $(el).parent().remove()
}

function getAllFilterWords() {
    return $("." + filterClass).map(function () {
        return $(this).val();
    }).get();
}