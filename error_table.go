
/* GENERATED FILE, DO NOT EDIT */
package log

type H map[string]string

var ErrorTable = map[E]H{
	"invalid_kb_handle": H{
		"vi_VN": "Địa chỉ không hợp lệ, vui lòng chọn địa chỉ dài hơn 5 ký tự và chỉ chứa ký tự chữ cái latin, chữ số hoặc '-', '_'",
		"en_US": "Invalid handle, please select a handle longer than 5 characters and containing only latin alphabetic characters, digits or '-', '_'",
	},
	"unsupported_file_format": H{
		"vi_VN": "Định dạng file {file_format} không được hỗ trợ. Danh sách những định dạng file được hỗ trợ: {supported}.",
		"en_US": "The file format {file_format} is not supported. List of supported file formats: {supported}.",
	},
	"api_http_unsupported": H{
		"vi_VN": "API chỉ có thể truy cập qua HTTPS. Đảm bảo URL bắt đầu bằng 'https://' chứ không phải 'http://'.",
		"en_US": "The API is only accessible over HTTPS. Ensure the URL starts with 'https://' and not 'http://'.",
	},
	"kb_handle_already_used": H{
		"vi_VN": "Địa chỉ đã được sử dụng, vui lòng chọn địa chỉ khác",
		"en_US": "Handle already in use, please choose another handle",
	},
	"still_have_open_invoice": H{
		"vi_VN": "Không thể thực hiện vì tài khoản của bạn vẫn còn hóa đơn chưa thanh toán. Vui lòng thử lại sau khi thanh toán",
		"en_US": "Unable to proceed because your account still has an open invoice. Please try again after completing the payment.",
	},
	"network_error": H{
		"vi_VN": "Không thể kết nối với máy chủ do sự cố mạng, vui lòng kiểm tra kết nối internet của bạn sau đó thử lại.",
		"en_US": "Cannot connect to the server due to network problem, please check your internet connection then try again.",
	},
	"invalid_subscription": H{
		"vi_VN": "Trạng thái gói cước không hợp lệ. Vui lòng liên hệ Subiz để nhận hỗ trợ",
		"en_US": "Subscription is invalid. Please contact Subiz for support",
	},
	"password_too_weak": H{
		"vi_VN": "Người dùng này đã bị chặn, vui lòng bỏ chặn để tiếp tục",
		"en_US": "This user have been banned, please unban this user to continue",
	},
	"close_public_channel": H{
		"vi_VN": "Không thể đóng hội thoại ở kênh công khai",
		"en_US": "Cannot end conversation in this public channel",
	},
	"invalid_promotion_code": H{
		"vi_VN": "Mã khuyến mại không hợp lệ. Vui lòng thử lại với mã khác",
		"en_US": "Promotion code is not valid. Please try again with another code",
	},
	"fb_outside_send_window": H{
		"vi_VN": "Không thể gửi tin ngoài thời gian cho phép. Bạn chỉ có thể nhắn tin khi khách chủ động nhắn tin trước. Đây là chính sách chung của facebook, chi tiết https://developers.facebook.com/docs/messenger-platform/policy/policy-overview",
		"en_US": "Cannot send messages outside the allowed time. You can only text when the customer proactively texts first. This is facebook's general policy, details https://developers.facebook.com/docs/messenger-platform/policy/policy-overview",
	},
	"google_error": H{
		"vi_VN": "Gián đoạn kết nối tới máy chủ Google, vui lòng thử lại sau",
		"en_US": "Connection to Google server interrupted, please try again later",
	},
	"invalid_otp": H{
		"vi_VN": "Mã OTP đã hết hạn",
		"en_US": "This OTP code has expired",
	},
	"invalid_webhook_url": H{
		"vi_VN": "Địa chỉ webhook không hợp lệ, vui lòng sử dụng một tên miền công khai",
		"en_US": "Invalid webhook URL, please use a public domain",
	},
	"invalid_integration": H{
		"vi_VN": "Cài đặt kênh giao tiếp bị ngắt hoặc không tồn tại. Vui lòng cài dặt lại kênh giao tiếp",
		"en_US": "Communication channel settings are disconnected or do not exist. Please reset the communication channel",
	},
	"invalid_currency": H{
		"vi_VN": "Đơn vị tiền tệ không hợp lệ. Vui lòng chọn USD hoặc VND",
		"en_US": "Invalid currency. Please select USD or VND",
	},
	"inactive_number": H{
		"vi_VN": "Đầu số {number} chưa được kích hoạt hoặc đã bị khóa, vui lòng kích hoạt đầu số hoặc kiểm tra lại cài đặt",
		"en_US": "The number {number} is not activated or locked, please activate the number or check the settings again",
	},
	"invite_link_expired": H{
		"vi_VN": "Link mời đã hết hạn. Vui lòng liên hệ chủ tài khoản để lấy link mới",
		"en_US": "The invitation link has expired. Please contact the account owner to get a new link",
	},
	"blocked_number": H{
		"vi_VN": "Số điện thoại {number} đang bị chặn. Vui lòng mở chặn để tiếp tục",
		"en_US": "Phone number {number} is blocked. Please unblock to continue",
	},
	"invalid_password_length": H{
		"vi_VN": "Mật khẩu quá ngắn, vui lòng chọn mật khẩu nhiều hơn {required_length} ký tự",
		"en_US": "Password is too short, please choose a password with more than {required_length} characters",
	},
	"conversation_ended": H{
		"vi_VN": "Hội thoại đã kết thúc, bạn không thể thực hiện hành động trọng hội thoại này",
		"en_US": "The conversation has ended, you cannot perform actions in this conversation",
	},
	"invalid_token": H{
		"vi_VN": "Mã không hợp lệ",
		"en_US": "The token is not valid",
	},
	"malformed_request": H{
		"vi_VN": "Yêu cầu không đúng. Vui lòng liên hệ Subiz để được hỗ trợ",
		"en_US": "Malformed request. Please contact Subiz for support",
	},
	"email_taken": H{
		"vi_VN": "Email {email} đã được sử dụng, vui lòng sử dụng email khác",
		"en_US": "Email {email} is already taken, please use another email",
	},
	"invalid_connection": H{
		"vi_VN": "Kết nối không hợp lệ, vui lòng kết nối lại",
		"en_US": "Your connection is invalid, please reconnect",
	},
	"invalid_poll_connection": H{
		"vi_VN": "Kết nối thời gian thực không hợp lệ, vui lòng kết nối lại",
		"en_US": "Real-time connection is invalid, please reconnect",
	},
	"dead_poll_connection": H{
		"vi_VN": "Kết nối thời gian thực đã bị ngắt do nghẽn mạng, vui lòng thử lại sau",
		"en_US": "Real-time connection was interrupted due to network congestion, please try again later",
	},
	"duplicate_contact": H{
		"vi_VN": "Giá trị {prop} đã được sử dụng cho một hồ sơ khách khác. Để có thể cập nhật, bạn cần gộp 2 hồ sơ hoặc xóa giá trị này ở hồ sơ còn lại.",
		"en_US": "The property {prop} is already in used for another contact. To continue, please merge 2 contacts or remove this value in the other contact.",
	},
	"insufficient_credit": H{
		"vi_VN": "Tài khoản của bạn không còn đủ credit. Vui lòng thanh toán tài khoản '{credit_name}' để tiếp tục",
		"en_US": "Your account has insufficient credits. Please pay for billing account '{credit_name}' to continue",
	},
	"invalid_facebook_token": H{
		"vi_VN": "Kết nối tới Facebook đã hết hạn. Vui lòng tích hợp lại fanpage {page_name} ({page_id}) để tiếp tục",
		"en_US": "Connection to Facebook has expired. Please reintegrate fanpage {page_name} ({page_id}) to continue",
	},
	"invalid_field": H{
		"vi_VN": "Trường dữ liệu {name} không hợp lệ. Vui lòng thử lại với giá trị khác",
		"en_US": "The field {name} is not valid. Please try again with a different value",
	},
	"invalid_zalo_token": H{
		"vi_VN": "Kết nối tới Zalo đã hết hạn. Vui lòng tích hợp lại OA {oa_name} ({oa_id}) để tiếp tục",
		"en_US": "Connection to Zalo has expired. Please reintegrate OA {oa_name} ({oa_id}) to continue",
	},
	"invalid_google_token": H{
		"vi_VN": "Kết nối tới tài khoản Google đã hết hạn. Vui lòng tích hợp lại địa điểm kinh doanh {location_name} ({location_id}) để tiếp tục",
		"en_US": "Connection to Google Account has expired. Please reintegrate your business location {location_name} ({location_id}) to continue",
	},
	"service_unavailable": H{
		"vi_VN": "Không thể kết nối tới dịch vụ cần thiết. Vui lòng thử lại sau",
		"en_US": "Unable to connect to the required service. Please try again later",
	},
	"payload_too_large": H{
		"vi_VN": "Dung lượng gói tin quá lớn",
		"en_US": "Payload too large",
	},
	"limit_exceeded": H{
		"vi_VN": "Bạn đã sử dụng quá giới hạn cho phép ({capacity})",
		"en_US": "You have exceeded the allowable limit ({capacity})",
	},
	"rate_limit": H{
		"vi_VN": "Yêu cầu bị từ chối do vượt quá giới hạn cho phép. Bạn cần phải chậm lại",
		"en_US": "Request denied due to exceeding allowed limit. You need to slow down",
	},
	"invalid_domain": H{
		"vi_VN": "",
		"en_US": "Something wrong, please try again later",
	},
	"missing_id": H{
		"vi_VN": "Lỗi định danh {type} không hợp lệ. Vui lòng cung cấp đầy đủ định danh hoặc liên hệ Subiz để được hỗ trợ",
		"en_US": "Invalid identify for {type}. Please provide the corrected identify or contact Subiz for support",
	},
	"leaver_is_the_last_one_in_conversation": H{
		"vi_VN": "Bạn không thể rời hội thoại, bạn là agent cuối cùng trong hội thoại này. Vui lòng thêm agent khác trước khi rời",
		"en_US": "You cannot leave this conversation, you are the last agent in the conversation. Please add another agent before leaving",
	},
	"not_a_conversation_member": H{
		"vi_VN": "Bạn không phải là thành viên của hội thoại. Bạn cần được mời vào hội thoại để tiếp tục",
		"en_US": "You are not a member of this conversation. You need to be invited before continue this action",
	},
	"transform_data": H{
		"vi_VN": "",
		"en_US": "Something wrong, please try again later",
	},
	"provider_failed": H{
		"vi_VN": "Yêu cầu thất bại từ {external_service}. Vui lòng thử lại sau",
		"en_US": "Your request to {external_service} failed. Please try again later",
	},
	"locked_account": H{
		"vi_VN": "Tài khoản của bạn đang bị khóa. Vui lòng liên hệ chủ tài khoản hoặc Subiz để được hỗ trợ",
		"en_US": "Your account is locked. Please contact account owner or Subiz for support",
	},
	"locked_agent": H{
		"vi_VN": "Tài khoản của bạn đang bị khóa. Vui lòng liên hệ chủ tài khoản hoặc Subiz để được hỗ trợ",
		"en_US": "Your account is locked. Please contact account owner or Subiz for support",
	},
	"provider_data_mismatched": H{
		"vi_VN": "Bất đồng bộ dữ liệu với {external_service}.",
		"en_US": "Data type mismatch with {external_service}.",
	},
	"file_system_error": H{
		"vi_VN": "Lỗi hệ thống tệp. Vui lòng thử lại hoặc liên hệ Subiz để được hỗ trợ",
		"en_US": "File system error. Please retry or contact Subiz for support",
	},
	"access_token_expired": H{
		"vi_VN": "Phiên đăng nhập đã kết thúc. Vui lòng đăng nhập lại hoặc thử lại với một token khác hợp lệ.",
		"en_US": "The login session has ended. Please log in again or try again with a different valid token.",
	},
	"internal_connection": H{
		"vi_VN": "Lỗi kết nối nội bộ, vui lòng thử lại sau",
		"en_US": "Internal connection error, please retry later",
	},
	"missing_resource": H{
		"vi_VN": "Không tìm thấy {type}",
		"en_US": "{type} not found",
	},
	"not_found": H{
		"vi_VN": "Không tìm thấy {type}",
		"en_US": "{type} not found",
	},
	"internal": H{
		"vi_VN": "Lỗi hệ thống. Vui lòng thử lại sau",
		"en_US": "System error. Please try again later",
	},
	"access_deny": H{
		"vi_VN": "Từ chối truy cập",
		"en_US": "Access deny",
	},
	"locked_user": H{
		"vi_VN": "Tài khoản của bạn đã bị khóa",
		"en_US": "Your account had been locked",
	},
	"unauthorized": H{
		"vi_VN": "Bạn không có đủ quyền. Để thực hiện chức năng này bạn cần thêm quyền {mising}",
		"en_US": "You are not authorized. To perform this action you need following permission: {missing}",
	},
	"invalid_input": H{
		"vi_VN": "Dữ liệu đầu vào không hợp lệ",
		"en_US": "Invalid input data",
	},
	"invalid_input_format": H{
		"vi_VN": "Trường dữ liệu '{invalid_field}' không hợp lệ. Vui lòng thử với giá trị khác ({msg}).",
		"en_US": "The data field '{invalid_field}' is not valid. Please try with another value ({msg}).",
	},
	"invalid_email": H{
		"vi_VN": "Email không hợp lệ",
		"en_US": "Email is not valid",
	},
	"weak_password": H{
		"vi_VN": "Mật khẩu quá yếu. Mật khẩu phải chứa ít nhất 8 ký tự",
		"en_US": "Password is too week. Password should contains at least 8 chracters",
	},
	"user_is_banned": H{
		"vi_VN": "Người dùng này đã bị chặn, vui lòng bỏ chặn để tiếp tục",
		"en_US": "This user have been banned, please unban this user to continue",
	},
	"user_is_unsubscribed": H{
		"vi_VN": "người dùng này đã từ chối nhận tin truyền thông của bạn, vui lòng bỏ người dùng khỏi danh sách từ chối nhận để tiếp tục",
		"en_US": "This user have unsubscribed your marketing messages, please remove this usser from unsubscribe list to continue",
	},
	"agent_is_inactived": H{
		"vi_VN": "Người dùng này đã bị chặn, vui lòng bỏ chặn để tiếp tục",
		"en_US": "This user have been banned, please unban this user to continue",
	},
	"wrong_password": H{
		"vi_VN": "Sai mật khẩu",
		"en_US": "Wrong password",
	},
}
