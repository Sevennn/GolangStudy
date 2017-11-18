var scope = 4;
$(".sub-btn").click(function() {
    let data = { 'firstname': $('#firstName').val(), 'lastname': $('#lastName').val(), 'username': $('#userName').val() };

    $.post('/update/userinfo', data, function(res) {
        if (res != undefined) {
            $('.table tbody').append(
                `<tr>
                  <th scope="row">${scope}</th>
                  <td>${res.firstName}</td>
                  <td>${res.lastName}</td>
                  <td>${res.userName}</td>
                </tr>`)
            scope++;
        }
    })

    console.log(data);
})