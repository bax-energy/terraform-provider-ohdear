package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceStatusPage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStatusPageCreate,
		ReadContext:   resourceStatusPageRead,
		UpdateContext: resourceStatusPageUpdate,
		DeleteContext: resourceStatusPageDelete,

		Schema: map[string]*schema.Schema{
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL of the site to be monitored.",
			},
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The ID of the team that owns the site.",
			},
			"sites": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of checks to be performed on the site.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "id of the site you want to add",
						},
						"clickable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Check the performance of the site.",
						},
					},
				},
			},
		},
		CustomizeDiff: resourceOhdearStatusPageDiff,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceOhdearStatusPageDiff(_ context.Context, d *schema.ResourceDiff, m interface{}) error {
	// set team_id from provider default if not provided
	if v, ok := d.GetOk("team_id"); !ok || v.(string) == "" {
		if err := d.SetNew("team_id", m.(*Config).teamID); err != nil {
			return err
		}
	}

	return nil
}
func resourceStatusPageCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Config).client

	// Costruisci il payload per la creazione
	payload := BuildStatusPage(d, "create")

	// Invia il payload al client per creare la status page
	statusPage, err := client.AddStatusPage(payload)
	if err != nil {
		return diagErrorf(err, "Could not add status page to Oh Dear")
	}

	// Imposta l'ID della risorsa nello stato Terraform
	d.SetId(fmt.Sprintf("%d", statusPage.ID))

	// Leggi i dati della risorsa appena creata per sincronizzarli con lo stato Terraform
	return resourceStatusPageRead(ctx, d, m)
}

func resourceStatusPageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implementa la logica di lettura della risorsa qui
	return diag.Diagnostics{}
}

func resourceStatusPageUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Config).client

	if d.HasChange("sites") {
		// Get the current state of sites
		old, new := d.GetChange("sites")
		oldSites := old.([]interface{})
		newSites := new.([]interface{})

		// Remove all existing sites
		for _, site := range oldSites {
			siteData := site.(map[string]interface{})
			siteID, _ := strconv.Atoi(siteData["id"].(string))
			err := client.RemoveSiteStatusPage(d.Id(), siteID)
			if err != nil {
				return diagErrorf(err, "Could not remove site %d from status page", siteID)
			}
		}

		// Re-add the updated list of sites
		updatedSites := []map[string]interface{}{}
		for _, site := range newSites {
			siteData := site.(map[string]interface{})
			siteID, _ := strconv.Atoi(siteData["id"].(string))
			updatedSites = append(updatedSites, map[string]interface{}{
				"id":        siteID,
				"clickable": siteData["clickable"].(bool),
			})
		}

		if len(updatedSites) > 0 {
			err := client.AddSiteStatusPage(d.Id(), map[string]interface{}{
				"sync":  true,
				"sites": updatedSites,
			})
			if err != nil {
				return diagErrorf(err, "Failed to sync updated sites to status page")
			}
		}
	}

	return resourceStatusPageRead(ctx, d, m)
}

func getStatusPageID(d *schema.ResourceData) (int, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return id, fmt.Errorf("corrupted resource ID in terraform state, Oh Dear only supports integer IDs. Err: %w", err)
	}
	return id, err
}

func resourceStatusPageDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id, err := getStatusPageID(d)
	if err != nil {
		return diag.FromErr(err)
	}

	client := m.(*Config).client
	if err = client.RemoveStatusPage(id); err != nil {
		return diagErrorf(err, "Could not remove site %d from Oh Dear", id)
	}

	return nil
}

func BuildStatusPage(d *schema.ResourceData, action string) map[string]interface{} {
	payload := map[string]interface{}{
		"title":   d.Get("title").(string),
		"team_id": d.Get("team_id").(string),
	}

	// Build the list of sites
	sites := d.Get("sites").([]interface{})
	siteList := make([]map[string]interface{}, len(sites))
	for i, site := range sites {
		siteData := site.(map[string]interface{})
		siteID, _ := strconv.Atoi(siteData["id"].(string)) // Convert id to integer
		siteList[i] = map[string]interface{}{
			"id":        siteID,
			"clickable": siteData["clickable"].(bool),
		}
	}
	payload["sites"] = siteList

	return payload
}
