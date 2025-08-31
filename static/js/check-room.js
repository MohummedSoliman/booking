export function checkRoomAvailability(roomId) {
  document
    .getElementById("check-availability-btn")
    .addEventListener("click", () => {
      let html = `
    <form id="check-availability-form" action="post" action="" novalidate class="needs-validation">
      <div class="row">
        <div class="col">
          <div class="row" id="reservation-dates-model">
            <div class="col">
              <input disabled required class="form-control" type="text" name="start" id="start" placeholder="Arrival"/>
            </div>
            <div class="col">
              <input disabled required class="form-control" type="text" name="end" id="end" placeholder="Depature"/>
            </div>
          </div>
        </div>
      </div>
    </form>
  `;
      attention.custom({
        title: "Choose Your Dates",
        msg: html,
        willOpen: () => {
          const elem = document.getElementById("reservation-dates-model");
          const rangePicker = new DateRangePicker(elem, {
            format: "yyyy-mm-dd",
            showOnFoucs: true,
            minDate: new Date(),
          });
        },
        didOpen: () => {
          document.getElementById("start").removeAttribute("disabled");
          document.getElementById("end").removeAttribute("disabled");
        },
        callback: function (result) {
          let form = document.getElementById("check-availability-form");

          let formData = new FormData(form);
          formData.append("csrf_token", "{{.CSRFToken}}");
          formData.append("room_id", roomId);

          formData.forEach((k, v) => console.log(k, v));
          console.log("//////////////////////////");
          fetch("/search-availability-json", {
            method: "POST",
            body: formData,
          })
            .then((res) => res.json())
            .then((data) => {
              if (data.ok) {
                attention.custom({
                  icon: "success",
                  msg: `<p>Room is Available</p>
                                    <p><a
                                    href="/book-room?id=${data.room_id}&sd=${data.start_date}&ed=${data.end_date}"
                                          class="btn btn-primary">Book Now!
                                    </a></p>`,
                  showConfirmButton: false,
                });
              } else {
                attention.error({
                  msg: "No Availability",
                });
              }
            });
        },
      });
    });
}
