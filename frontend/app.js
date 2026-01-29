const form = document.getElementById("messageForm");
const btn = document.getElementById("sendBtn");

// PAKAI RELATIVE URL
const API_URL = "/send";

// kalau sudah deploy:
// const API_URL = "https://nama-backend.up.railway.app/send";

form.addEventListener("submit", async (e) => {
  e.preventDefault();

  const phone = document.getElementById("phone").value.trim();
  const message = document.getElementById("message").value.trim();

  if (!phone || !message) {
    Swal.fire("Oops", "Nomor dan pesan wajib diisi", "warning");
    return;
  }

  btn.disabled = true;
  btn.innerText = "Mengirim...";

  try {
    const res = await fetch(API_URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        phone: phone,
        message: message,
      }),
    });

    if (!res.ok) {
      throw new Error("Gagal mengirim pesan");
    }

    Swal.fire({
      icon: "success",
      title: "Berhasil",
      text: "Pesan anonim berhasil dikirim",
      confirmButtonText: "OK",
    });

    form.reset();
  } catch (err) {
    Swal.fire({
      icon: "error",
      title: "Gagal",
      text: err.message || "Terjadi kesalahan",
    });
  } finally {
    btn.disabled = false;
    btn.innerText = "Kirim Pesan";
  }
});

