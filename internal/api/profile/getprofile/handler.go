package getprofile

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gameap/gameap/internal/api/base"
	"github.com/gameap/gameap/internal/domain"
	"github.com/gameap/gameap/internal/repositories"
	"github.com/gameap/gameap/internal/services/mfanudge"
	"github.com/gameap/gameap/pkg/api"
	"github.com/gameap/gameap/pkg/auth"
	"github.com/pkg/errors"
)

type Handler struct {
	rbac      repositories.RBACRepository
	nudge     *mfanudge.Service
	admin     adminChecker
	responder base.Responder
}

func NewHandler(
	rbac repositories.RBACRepository,
	nudge *mfanudge.Service,
	admin adminChecker,
	responder base.Responder,
) *Handler {
	return &Handler{
		rbac:      rbac,
		nudge:     nudge,
		admin:     admin,
		responder: responder,
	}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session := auth.SessionFromContext(ctx)
	if !session.IsAuthenticated() {
		h.responder.WriteError(ctx, rw, api.WrapHTTPError(
			errors.New("user not authenticated"),
			http.StatusUnauthorized,
		))

		return
	}

	roles, err := h.rbac.GetRolesForEntity(ctx, session.User.ID, domain.EntityTypeUser)
	if err != nil {
		h.responder.WriteError(ctx, rw, errors.WithMessage(err, "failed to get user roles"))

		return
	}

	h.responder.Write(ctx, rw, newProfileResponseFromUser(session.User, roles, h.mfaNudge(ctx, session.User)))
}

// mfaNudge computes the admin-MFA recommendation for the current user, or nil
// when none applies. The profile read is side-effect-free: unlike the login
// flow it does not persist the first-shown timestamp (login owns that write).
func (h *Handler) mfaNudge(ctx context.Context, user *domain.User) *mfanudge.View {
	if h.nudge == nil || h.admin == nil {
		return nil
	}

	isAdmin, err := h.admin.Can(ctx, user.ID, []domain.AbilityName{domain.AbilityNameAdminRolesPermissions})
	if err != nil {
		slog.WarnContext(ctx, "failed to check admin status for MFA nudge", "error", err)

		return nil
	}

	return mfanudge.NewView(h.nudge.Compute(mfanudge.Snapshot{
		IsAdmin:          isAdmin,
		TwoFactorEnabled: user.TwoFactorEnabled,
		FirstShownAt:     user.MFAFirstShownAt(),
		SnoozedUntil:     user.MFASnoozedUntil(),
	}))
}
