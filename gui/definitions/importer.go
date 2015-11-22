
package definitions

func init(){
  add(`Importer`, &defImporter{})
}

type defImporter struct{}

func (*defImporter) String() string {
	return `
<?xml version="1.0" encoding="utf-8"?>
<interface>
  <object class="GtkListStore" id="importAccountsStore">
    <columns>
      <column type="gchararray"/>
      <column type="gchararray"/>
      <column type="gboolean"/>
    </columns>
  </object>
  <object class="GtkDialog" id="importerWindow">
    <property name="window-position">1</property>
    <property name="title">$title</property>
    <property name="width_request">450</property>
    <property name="height_request">600</property>
    <property name="border_width">10</property>
    <child internal-child="vbox">
      <object class="GtkVBox" id="box">
        <property name="homogeneous">false</property>
        <child>
          <object class="GtkLabel" id="label">
	        <property name="label">$importMessage</property>
            <property name="wrap">true</property>
            <property name="wrap-mode">2</property>
	      </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkScrolledWindow" id="importerScroll">
            <property name="vscrollbar-policy">GTK_POLICY_AUTOMATIC</property>
            <property name="hscrollbar-policy">GTK_POLICY_AUTOMATIC</property>
            <child>
              <object class="GtkTreeView" id="importerTreeView">
                <property name="model">importAccountsStore</property>
                <child>
                  <object class="GtkTreeViewColumn" id="import-this-account-column">
                    <property name="title">$columnTitleImportThis</property>
                    <property name="sort_column_id">2</property>
                    <child>
                      <object class="GtkCellRendererToggle" id="import-this-account-renderer">
                        <property name="activatable">true</property>
                      </object>
                      <attributes>
                        <attribute name="active">2</attribute>
                      </attributes>
                    </child>
                  </object>
                </child>
                <child>
                  <object class="GtkTreeViewColumn" id="from-application-column">
                    <property name="title">$columnTitleFromApplication</property>
                    <property name="sort_column_id">0</property>
                    <child>
                      <object class="GtkCellRendererText" id="from-application-renderer"/>
                      <attributes>
                        <attribute name="text">0</attribute>
                      </attributes>
                    </child>
                  </object>
                </child>
                <child>
                  <object class="GtkTreeViewColumn" id="account-name-column">
                    <property name="title">$columnTitleAccountName</property>
                    <property name="sort_column_id">1</property>
                    <child>
                      <object class="GtkCellRendererText" id="account-name-renderer"/>
                      <attributes>
                        <attribute name="text">1</attribute>
                      </attributes>
                    </child>
                  </object>
                </child>
              </object>
            </child>
          </object>
          <packing>
            <property name="expand">true</property>
            <property name="fill">true</property>
            <property name="position">1</property>
          </packing>
        </child>
      </object>
    </child>
    <child type="action">
      <object class="GtkButton" id="button_cancel">
        <property name="label">$buttonTitleCancel</property>
      </object>
    </child>
    <child type="action">
      <object class="GtkButton" id="button_ok">
        <property name="label">$buttonTitleImport</property>
        <property name="can-default">true</property>
      </object>
    </child>
    <action-widgets>
      <action-widget response="cancel">button_cancel</action-widget>
      <action-widget response="ok" default="true">button_ok</action-widget>
    </action-widgets>
  </object>
</interface>

`
}