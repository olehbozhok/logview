function showModalLog(selector) {
    $("#json").html($(selector).html());

    $('#myModal').modal('show')
}

// var json = {
//     "hey": "guy",
//     "anumber": 243,
//     "anobject": {
//         "whoa": "nuts",
//         "anarray": [1, 2, "thr<h1>ee"],
//         "more": "stuff"
//     },
//     "awesome": true,
//     "bogus": false,
//     "meaning": null,
//     "japanese": "明日がある。",
//     "link": "http://jsonview.com",
//     "notLink": "http://jsonview.com is great",
//     "multiline": ['Much like me, you make your way forward,',
//         'Walking with downturned eyes.',
//         'Well, I too kept mine lowered.',
//         'Passer-by, stop here, please.'].join("\n")
// };

// TODO: add view json via JSONView
// $(document).ready(function () {

//     $(function () {
//         // $("#json").JSONView(json);

//         // $("#json-collapsed").JSONView(json, { collapsed: true, nl2br: true, recursive_collapser: true });

//         $('#collapse-btn').on('click', function () {
//             $('#json').JSONView('collapse');
//         });

//         $('#expand-btn').on('click', function () {
//             $('#json').JSONView('expand');
//         });

//         $('#toggle-btn').on('click', function () {
//             $('#json').JSONView('toggle');
//         });

//         $('#toggle-level1-btn').on('click', function () {
//             $('#json').JSONView('toggle', 1);
//         });

//         $('#toggle-level2-btn').on('click', function () {
//             $('#json').JSONView('toggle', 2);
//         });
//     });
// })